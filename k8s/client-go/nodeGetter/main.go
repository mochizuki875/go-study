/*
 client-goでNodeの情報を取得する
*/

package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

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

	// get clientset for kubernetes resources
	// clientsetは各API Groupのリソースを操作するためのclientの集まり
	clientset, _ := kubernetes.NewForConfig(config)

	// clientsetを使ってNodeを取得
	// Get Node objects
	node, _ := clientset.CoreV1().Nodes().Get(ctx, "k8s-cluster2-node01", metav1.GetOptions{})

	// status.conditions[?(type=Ready)]を取得
	fmt.Printf("Conditions: %#v\n", getNodeConditionReady(node))

	// nodeがRedyならtrue
	fmt.Println("Ready: ", checkNodeReady(node))

}

func getNodeConditionReady(node *v1.Node) v1.NodeCondition {
	var nodeConditionReady v1.NodeCondition
	for _, nodeCondition := range node.Status.Conditions {
		if nodeCondition.Type == v1.NodeReady {
			nodeConditionReady = nodeCondition
		}
	}
	return nodeConditionReady
}

func checkNodeReady(node *v1.Node) bool {
	for _, condition := range node.Status.Conditions {
		if condition.Type == v1.NodeReady && condition.Status == v1.ConditionTrue {
			return true
		}
	}
	return false
}
