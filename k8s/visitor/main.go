/*
 visitorを使ってPod内の全てのコンテナにアクセスする
*/

package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/kubectl/pkg/util/podutils"
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

	// clientsetを使ってPodを取得
	// Get pod objects
	pod, _ := clientset.CoreV1().Pods("default").Get(ctx, "test-pod-visitor", metav1.GetOptions{})

	// Podに含まれる各種Containerについてvisitorを利用してコンテナ名を表示
	printContainerName(pod, podutils.InitContainers)
	printContainerName(pod, podutils.Containers)
	printContainerName(pod, podutils.EphemeralContainers)
}

// visitを利用してコンテナ名を表示
// visitorとしてコンテナ名を表示する関数を渡している
func printContainerName(pod *corev1.Pod, containerType podutils.ContainerType) {
	// 第3引数のvisitor関数でfaulseを返すと処理中断
	podutils.VisitContainers(&pod.Spec, containerType, func(c *corev1.Container, _ podutils.ContainerType) bool {
		switch {
		case containerType == podutils.InitContainers:
			fmt.Println("InitContainer: ", c.Name)
		case containerType == podutils.EphemeralContainers:
			fmt.Println("EphemeralContainer: ", c.Name)
		default:
			fmt.Println("Container: ", c.Name)
		}
		return true
	})
}
