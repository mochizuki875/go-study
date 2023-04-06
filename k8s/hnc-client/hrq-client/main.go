/*
  Namespace配下のHierarchicalResourceQuotasを取得しリ設定されたソースに対する使用状況を表示する(ResourceQuotasのイメージ)
  ResourceQuotasはAPI Serverで表示項目を整形するロジックを持っているがHierarchicalResourceQuotasは持っていないかつCRD定義でも実現不可なので
  クライアントサイドに本来API Serverで実現する処理を追加
*/

package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/util/duration"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	"example.com/hrq-client/api/v1alpha2"

	"github.com/liggitt/tabwriter"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sprinters "k8s.io/kubernetes/pkg/printers"
)

var namespace = "test"
// var names = []string{}
var names = []string{"test-hrq"}
// var names = []string{"test-hrq", "test-hrq-2"}

var hncClient *rest.RESTClient

// tabwriterでテーブル表示する際のOP
// 以下から流用
// https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/cli-runtime/pkg/printers/tabwriter.go
const (
	tabwriterMinWidth = 6
	tabwriterWidth    = 4
	tabwriterPadding  = 3
	tabwriterPadChar  = ' '
	tabwriterFlags    = tabwriter.RememberWidths
)

// Column定義(ColumnDefinitions)
// kube-apiserverで用いられている以下を流用
// https://github.com/kubernetes/kubernetes/blob/f28e9f6f45974909968216cf4cc8872b72170335/pkg/printers/internalversion/printers.go#L455-L460
var hierarchicalResourceQuotaColumnDefinitions = []metav1.TableColumnDefinition{
	{Name: "Name", Type: "string", Format: "name", Description: metav1.ObjectMeta{}.SwaggerDoc()["name"]},
	{Name: "Age", Type: "string", Description: metav1.ObjectMeta{}.SwaggerDoc()["creationTimestamp"]},
	{Name: "Request", Type: "string", Description: "Request represents a minimum amount of cpu/memory that a container may consume."},
	{Name: "Limit", Type: "string", Description: "Limits control the maximum amount of cpu/memory that a container may use independent of contention on the node."},
}

func main() {
	//////////////////////////////////////////////事前処理/////////////////////////////////////////////////////////
	var defaultKubeConfigPath string

	// kubeconfigのパスを文字列として作成
	if home := homedir.HomeDir(); home != "" {
		// build kubeconfig path from $HOME dir
		defaultKubeConfigPath = filepath.Join(home, ".kube", "config")
	}

	// kubeconfigフラグから値を取得
	// set kubeconfig flag
	kubeconfig := flag.String("kubeconfig", defaultKubeConfigPath, "kubeconfig config file")
	flag.Parse()

	// create context
	ctx := context.Background()

	// clientset作成時に使うconfigを作成
	// retrieve kubeconfig
	config, _ := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	// create the HNC clientset
	hncConfig := *config
	hncConfig.ContentConfig.GroupVersion = &v1alpha2.GroupVersion
	hncConfig.APIPath = "/apis"
	hncConfig.NegotiatedSerializer = serializer.WithoutConversionCodecFactory{CodecFactory: scheme.Codecs}
	hncConfig.UserAgent = rest.DefaultKubernetesUserAgent()
	hncClient, _ := rest.UnversionedRESTClientFor(&hncConfig) // RestClientを作成(HNCのリソースはコレを使う)

	///////////////////////////////////////////////ここまで事前処理////////////////////////////////////////////////

	// Tableインスタンスの初期化
	table := &metav1.Table{ColumnDefinitions: hierarchicalResourceQuotaColumnDefinitions}


	option := k8sprinters.GenerateOptions{
		true,
		true,
	}

	// Namespace内のhierarchicalresourcequotasをListで取得
	hrqList := &v1alpha2.HierarchicalResourceQuotaList{}
	if len(names) > 0 { // 特定のリソース名を指定して取得
	    for _, name := range names {
	  	    hrq := &v1alpha2.HierarchicalResourceQuota{}
	  	    if err := hncClient.Get().Resource("hierarchicalresourcequotas").Namespace(namespace).Name(name).Do(ctx).Into(hrq); err != nil {
	  		    fmt.Printf("Error reading hierarchicalresourcequota %s: %s\n", name, err)
	  	    	os.Exit(1)
	   	    }
	        hrqList.Items = append(hrqList.Items, *hrq)
        }
    } else { // 全てのリソースを取得
		if err := hncClient.Get().Resource("hierarchicalresourcequotas").Namespace(namespace).Do(ctx).Into(hrqList); err != nil {
			fmt.Printf("Error reading hierarchicalresourcequotas: %s\n", err)
			os.Exit(1)
		}
	}

	if len(hrqList.Items) < 1 { // リソースが1件も見つからなければエラー文言を表示
		fmt.Printf("No resources found in %s namespace.\n", namespace)
		os.Exit(1)
	}


	tableRaws, err := printHierarchicalResourceQuotaList(hrqList, option) // 取得したListをtableRaws型に整形(本来サーバー側で実施)
	if err != nil {
		fmt.Printf("Error reading hierarchicalresourcequotas: %s\n", err)
	}
	table.Rows = tableRaws

	// witerの作成
	w := tabwriter.NewWriter(os.Stdout, tabwriterMinWidth, tabwriterWidth, tabwriterPadding, tabwriterPadChar, tabwriterFlags)

	// printer(クライアント側)を作成 ※printerはTablePrinterを使用
	p := printers.NewTablePrinter(printers.PrintOptions{
		NoHeaders:        false,
		WithNamespace:    true,
		WithKind:         true,
		Wide:             true,
		ShowLabels:       false,
		Kind:             schema.GroupKind{},
		ColumnLabels:     nil,
		SortBy:           "",
		AllowMissingKeys: false,
	})

	// printerを使ってwriterに書き込み
	p.PrintObj(table, w)

	// バッファ書き込み
	w.Flush() // ここで書き込んだ内容を出力

}

// 以下hierarchicalResourceQuotaのTableRowへの変換ロジック
// ResourceQuotaのものを流用
// https://github.com/kubernetes/kubernetes/blob/f28e9f6f45974909968216cf4cc8872b72170335/pkg/printers/internalversion/printers.go#L2536-L2578

// hierarchicalResourceQuotaから[]TableRowを作成
func printHierarchicalResourceQuota(hierarchicalResourceQuota *v1alpha2.HierarchicalResourceQuota, options k8sprinters.GenerateOptions) ([]metav1.TableRow, error) {
	row := metav1.TableRow{
		Object: runtime.RawExtension{Object: hierarchicalResourceQuota},
	}

	resources := make([]v1.ResourceName, 0, len(hierarchicalResourceQuota.Status.Hard))
	for resource := range hierarchicalResourceQuota.Status.Hard {
		resources = append(resources, resource)
	}
	sort.Sort(SortableResourceNames(resources))

	requestColumn := bytes.NewBuffer([]byte{})
	limitColumn := bytes.NewBuffer([]byte{})
	for i := range resources {
		w := requestColumn
		resource := resources[i]
		usedQuantity := hierarchicalResourceQuota.Status.Used[resource]
		hardQuantity := hierarchicalResourceQuota.Status.Hard[resource]

		// use limitColumn writer if a resource name prefixed with "limits" is found
		if pieces := strings.Split(resource.String(), "."); len(pieces) > 1 && pieces[0] == "limits" {
			w = limitColumn
		}

		fmt.Fprintf(w, "%s: %s/%s, ", resource, usedQuantity.String(), hardQuantity.String())
	}

	age := translateTimestampSince(hierarchicalResourceQuota.CreationTimestamp)
	row.Cells = append(row.Cells, hierarchicalResourceQuota.Name, age, strings.TrimSuffix(requestColumn.String(), ", "), strings.TrimSuffix(limitColumn.String(), ", "))
	return []metav1.TableRow{row}, nil
}

// hierarchicalResourceQuotaListから[]TableRowを作成
func printHierarchicalResourceQuotaList(list *v1alpha2.HierarchicalResourceQuotaList, options k8sprinters.GenerateOptions) ([]metav1.TableRow, error) {
	rows := make([]metav1.TableRow, 0, len(list.Items))
	for i := range list.Items {
		r, err := printHierarchicalResourceQuota(&list.Items[i], options)
		if err != nil {
			return nil, err
		}
		rows = append(rows, r...)
	}
	return rows, nil
}

// translateTimestampSince returns the elapsed time since timestamp in
// human-readable approximation.
func translateTimestampSince(timestamp metav1.Time) string {
	if timestamp.IsZero() {
		return "<unknown>"
	}

	return duration.HumanDuration(time.Since(timestamp.Time))
}

// SortableResourceNames - An array of sortable resource names
type SortableResourceNames []v1.ResourceName

func (list SortableResourceNames) Len() int {
	return len(list)
}

func (list SortableResourceNames) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list SortableResourceNames) Less(i, j int) bool {
	return list[i] < list[j]
}
