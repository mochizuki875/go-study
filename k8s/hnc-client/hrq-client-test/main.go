/*
  Namespace配下のhierarchicalresourcequotasを取得しリ設定されたソースに対する使用状況を表示する(resourcequotasのイメージ)
*/

package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	api "sigs.k8s.io/hierarchical-namespaces/api/v1alpha2"
)

var hncClient *rest.RESTClient

func main() {
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
	hncConfig.ContentConfig.GroupVersion = &api.GroupVersion
	hncConfig.APIPath = "/apis"
	hncConfig.NegotiatedSerializer = serializer.WithoutConversionCodecFactory{CodecFactory: scheme.Codecs}
	hncConfig.UserAgent = rest.DefaultKubernetesUserAgent()
	hncClient, _ := rest.UnversionedRESTClientFor(&hncConfig)

	// Namespace内のhierarchicalresourcequotasをListで取得
	hrq := &api.HierarchicalResourceQuotaList{}
	_ = hncClient.Get().Namespace("test").Name("hierarchicalresourcequotas").Do(ctx).Into(hrq)

	// 取得したhierarchicalresourcequotas単位でループ
	for i := range hrq.Items {

		// フィールドの中身を確認するためのデバッグ
		////////////////////
		// fmt.Printf("hrq_name: %#v\n", hrq.Items[i].Name)
		// fmt.Println("--------------------------------------------")
		// fmt.Printf("hrq_status_hard:\n %#v\n", hrq.Items[i].Status.Hard)
		// fmt.Println("--------------------------------------------")
		// fmt.Printf("hrq_status_used:\n %#v\n", hrq.Items[i].Status.Used)
		// fmt.Println("--------------------------------------------")
		// fmt.Printf("hrq.Status.Used.limits.cpu:\n %#v\n", hrq.Items[i].Status.Used["limits.cpu"])
		////////////////////

		Resources := []corev1.ResourceName{corev1.ResourceCPU, corev1.ResourceMemory, corev1.ResourceStorage, corev1.ResourceEphemeralStorage}

		// 結果を格納するsliceを作成
		var requests, limits []string

		for _, resource := range Resources { // 各リソースの使用状況を取得して結果をsliceに格納

			// requestsとlimitsのリソースフィールド名
			requests_resource := "requests." + resource
			limits_resource := "limits." + resource

			// Hard_Requests
			hard_requests := hrq.Items[i].Status.Hard[requests_resource]
			hard_requests_string := hard_requests.String()

			if hard_requests_string != "0" { // hard_requestsが定義されていれば使用状況を取得して結果をsliceに格納
				// Used
				used_requests := hrq.Items[i].Status.Used[requests_resource]
				used_requests_string := used_requests.String()

				// sliceに格納
				requests = append(requests, string(requests_resource)+": "+used_requests_string+"/"+hard_requests_string)

			}

			// Hard_Limits
			hard_limits := hrq.Items[i].Status.Hard[limits_resource]
			hard_limits_string := hard_limits.String()

			if hard_limits_string != "0" { // hard_limitsが定義されていれば使用状況を取得して結果をsliceに格納
				// Used
				used_limits := hrq.Items[i].Status.Used[limits_resource]
				used_limits_string := used_limits.String()

				// sliceに格納
				limits = append(limits, string(limits_resource)+": "+used_limits_string+"/"+hard_limits_string)

			}

		}

		fmt.Println("--------------------------------------------")

		// 結果出力
		fmt.Println("REQUEST: ", requests)
		fmt.Println("LIMIT: ", limits)
	}

}
