package main

import (
	"flag"
	"log"
	"path/filepath"
	"time"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var defaultKubeConfigPath string
	if home := homedir.HomeDir(); home != "" {
		// build kubeconfig path from $HOME dir
		defaultKubeConfigPath = filepath.Join(home, ".kube", "config")
	}

	// set kubeconfig flag
	kubeconfig := flag.String("kubeconfig", defaultKubeConfigPath, "kubeconfig config file")
	flag.Parse()

	config, _ := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// Create InformerFactory
	// 戻り値としてSharedInformerFactoryと言うinterfaceが指定されてるけど実際には初期化された構造体のポインタが返される
	// https://github.com/kubernetes/client-go/blob/v0.25.0/informers/factory.go#L96
	// https://github.com/kubernetes/client-go/blob/2666bd29867c96168c5e9429fd74fd9bfbedac8b/informers/factory.go#L109
	informerFactory := informers.NewSharedInformerFactory(clientset, time.Second*30)

	// Create pod informer by informerFactory
	podInformer := informerFactory.Core().V1().Pods()

	// Add EventHandler to informer
	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    func(new interface{}) { log.Println("Added") },
		UpdateFunc: func(old, new interface{}) { log.Println("Updated") },
		DeleteFunc: func(old interface{}) { log.Println("Deleted") },
	})

	// Start Go routines
	// InformerがPodを監視しin-memory-cacheにデータを格納
	informerFactory.Start(wait.NeverStop)
	// Wait until finish caching with List API
	informerFactory.WaitForCacheSync(wait.NeverStop)

	// Listerを作成してin-memory-cacheからPodをList
	// ================================================

	// Create Pod Lister
	podLister := podInformer.Lister()

	// Get List of pods
	// _, err = podLister.List(labels.Nothing())
	pods, err := podLister.List(labels.Everything()) // 全てのlabelを対象としてPodを取得
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Get List Success\n")
		for i, pod := range pods {
			log.Printf("[Pod Name %d]%s\n", i, pod.GetName())
		}
	}
	// ================================================

	for {
		time.Sleep(time.Second * 1)
	}
}
