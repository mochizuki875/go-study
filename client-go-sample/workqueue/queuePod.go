package main

import (
	"flag"
	"log"
	"path/filepath"
	"time"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/workqueue"
)

func main() {
	var defaultKubeConfigPath string

	// kubeconfigのデフォルトパス取得
	if home := homedir.HomeDir(); home != "" {
		// build kubeconfig path from $HOME dir
		defaultKubeConfigPath = filepath.Join(home, ".kube", "config")
	}

	// flagからkubeconfigを取得できるようにする(デフォルトはdefaultKubeConfigPath)
	kubeconfig := flag.String("kubeconfig", defaultKubeConfigPath, "kubeconfig config file")
	flag.Parse()

	// flagからkubeconfigをconfigとして取得
	config, _ := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	// kubeconfigを用いてclientsetを作成
	// https://pkg.go.dev/k8s.io/client-go@v0.25.0/kubernetes#Clientset
	// https://github.com/kubernetes/client-go/blob/v0.25.0/kubernetes/clientset.go#L128
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// 全Namespaceを対象としたInformerFactory(SharedInformerFactory)を作成
	// https://pkg.go.dev/k8s.io/client-go@v0.25.0/informers#SharedInformerFactory
	informerFactory := informers.NewSharedInformerFactory(clientset, time.Second*30)

	// InformerFactoryからPodInformerを作成
	// https://pkg.go.dev/k8s.io/client-go@v0.25.0/informers/core/v1#PodInformer
	podInformer := informerFactory.Core().V1().Pods()

	// Workqueueを作成
	// RateLimitingQueueとして作成
	// https://pkg.go.dev/k8s.io/client-go@v0.25.0/util/workqueue#NewRateLimitingQueue
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	// プロセス終了時にWorkqueueを終了する
	// https://pkg.go.dev/k8s.io/client-go@v0.25.0/util/workqueue#Interface
	defer queue.ShutDown()

	// PodInformerのEventHandlerにEventに対応した処理を追加
	// ResourceEventHandlerFuncsと言うstructが定義されている
	// https://pkg.go.dev/k8s.io/client-go@v0.25.0/tools/cache#ResourceEventHandlerFuncs
	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(old interface{}) {
			var key string
			var err error
			// in-memory-cacheからnamespace/nameを取得
			//https://pkg.go.dev/k8s.io/client-go@v0.25.0/tools/cache#MetaNamespaceKeyFunc
			if key, err = cache.MetaNamespaceKeyFunc(old); err != nil {
				runtime.HandleError(err)
				return
			}
			// Workqueueへのkey追加
			// https://pkg.go.dev/k8s.io/client-go@v0.25.0/util/workqueue#Type.Add
			queue.Add(key)
			log.Println("Added: " + key)
		},
		UpdateFunc: func(old, new interface{}) {
			var key string
			var err error
			if key, err = cache.MetaNamespaceKeyFunc(new); err != nil {
				runtime.HandleError(err)
				return
			}
			// Workqueueへのkey追加
			queue.Add(key)
			log.Println("Updated: " + key)
		},
		DeleteFunc: func(old interface{}) {
			var key string
			var err error
			if key, err = cache.MetaNamespaceKeyFunc(old); err != nil {
				runtime.HandleError(err)
				return
			}
			// Workqueueへのkey追加
			queue.Add(key)
			log.Println("Deleted: " + key)
		},
	})

	// informerを起動
	// https://pkg.go.dev/k8s.io/client-go/informers/internalinterfaces#SharedInformerFactory.Start
	// Start()の実装
	// https://github.com/kubernetes/client-go/blob/2666bd29867c96168c5e9429fd74fd9bfbedac8b/informers/factory.go#L128
	// informer単位でgo informer.Run(stopCh)と言う感じでgoroutineを起動している
	// wait.NeverStopはwaitパッケージ内で定義されており、戻り値は空のstruct型channel(chan struct{})
	informerFactory.Start(wait.NeverStop)

	// Wait until finish caching with List API
	// List APIでkube-apiserverから取得した値をin-memory-cacheに格納するまで待機
	// https://pkg.go.dev/k8s.io/client-go@v0.25.0/informers#SharedInformerFactory.WaitForCacheSync
	informerFactory.WaitForCacheSync(wait.NeverStop)

	// Listerを作成してin-memory-cacheからPodをList
	// ================================================

	// Create Pod Lister
	// in-memory-cacheからListするためのListerを作成
	podLister := podInformer.Lister()
	// Get List of pods
	// ListerからPodのListを取得
	// labels.Nothing()でラベルの付与されていないものを対象としている
	// labels.Everything()とすることで全てのラベルを対象
	// _, err = podLister.List(labels.Nothing())
	_, err = podLister.List(labels.Everything())

	if err != nil {
		log.Fatal(err)
	}
	// ================================================

	// 無限ループで処理を継続
	for {
		time.Sleep(time.Second * 1)
	}
}
