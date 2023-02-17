package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

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

	// clientsetを使ってPod一覧を取得(結果はPodListという構造体で変える、内部にItemsというPod型のsliceを持つ)
	// Get list of pod objects
	pods, _ := clientset.CoreV1().Pods("default").List(ctx, metav1.ListOptions{})

	// show pod object to stdout
	for i, pod := range pods.Items {
		fmt.Printf("[Pod Name %d]%s\n", i, pod.GetName())
	}
}
