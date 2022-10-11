/*
Workqueue Example
https://github.com/kubernetes/client-go/tree/master/examples/workqueue
*/

package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"time"

	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
)

type Controller struct {
	indexer  cache.Indexer                   // Indexer(in-memory-cacheとやりとりする人)
	queue    workqueue.RateLimitingInterface // WorkQueue
	informer cache.Controller                // Informer(kube-apiserverをList & Watch)
}

// コンストラクタ
func NewController(queue workqueue.RateLimitingInterface, indexer cache.Indexer, informer cache.Controller) *Controller {
	return &Controller{
		informer: informer,
		indexer:  indexer,
		queue:    queue,
	}
}

// Controllerとしてのプロセスを起動
func (c *Controller) Run(workers int, stopCh chan struct{}) {
	defer runtime.HandleCrash()
	defer c.queue.ShutDown()

	klog.Info("Start Pod Controller")

	go c.informer.Run(stopCh)

	if !cache.WaitForCacheSync(stopCh, c.informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Time out waiting for caches to sync"))
		return
	}

	for i := 0; i < workers; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh) // stopChがcloseするまでrunWorkerを1s間隔で実行し続ける
	}

	<-stopCh // ロック(stopChが受信するまでブロック)
	klog.Info("Stopping Pod controller")
}

// WorkerプロセスとしてprocessNextItemを実行し続ける
func (c *Controller) runWorker() {
	for c.processNextItem() {
	}
}

// WorkQueueから先頭のアイテムを取得
func (c *Controller) processNextItem() bool {
	key, quit := c.queue.Get() // WorkQueueから先頭のアイテムを取得
	if quit {
		return false
	}

	defer c.queue.Done(key)

	err := c.syncToStdout(key.(string)) // 処理に相当?

	c.HandleErr(err, key)

	return true
}

// 実処理に相当???(indexer経由でin-memory-cacheからオブジェクトを取得) Lister経由ではだめ？
func (c *Controller) syncToStdout(key string) error {
	obj, exists, err := c.indexer.GetByKey(key) // in-memory-cacheからアイテムを取得？
	if err != nil {
		klog.Errorf("Fetching object with key %s from store faild with %v", key, err)
		return err
	}

	if !exists {
		fmt.Printf("Pod %s dose not exist anymore\n", key)
	} else {
		fmt.Printf("Sync/Add/Update for Pod %s\n", obj.(*v1.Pod).GetName())
	}
	return nil
}

// 処理結果のハンドリング
func (c *Controller) HandleErr(err error, key interface{}) {
	if err == nil {
		c.queue.Forget(key)
		return
	}

	// 5回まではrequeueする
	if c.queue.NumRequeues(key) < 5 {
		klog.Infof("Error syncing pod %v: %v", key, err)
		klog.Infof("ReQueue %v", key)
		c.queue.AddRateLimited(key)
		return
	}

	c.queue.Forget(key)
	runtime.HandleError(err)
	klog.Infof("Dropping pod %q out of the queue: %v", key, err)
}

func main() {
	var defaultKubeConfigPath string

	// ローカルにあるkubeconfigを取得
	if home := homedir.HomeDir(); home != "" {
		defaultKubeConfigPath = filepath.Join(home, ".kube", "config")
	}

	var kubeconfig string
	var master string

	flag.StringVar(&kubeconfig, "kubeconfig", defaultKubeConfigPath, "absolute path to the kubeconfig file")
	flag.StringVar(&master, "master", "", "master url")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags(master, kubeconfig)
	if err != nil {
		klog.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatal(err)
	}

	podListWatcher := cache.NewListWatchFromClient(clientset.CoreV1().RESTClient(), "pods", v1.NamespaceDefault, fields.Everything())

	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	// indexerとinformerを作成
	indexer, informer := cache.NewIndexerInformer(podListWatcher, &v1.Pod{}, 0, cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(key)
			}
		},
		UpdateFunc: func(old interface{}, new interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(new)
			if err == nil {
				queue.Add(key)
			}
		},
		DeleteFunc: func(obj interface{}) {
			// IndexerInformer uses a delta queue, therefore for deletes we have to use this
			// key function.
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(key)
			}
		},
	}, cache.Indexers{})

	// Controllerを作成
	controller := NewController(queue, indexer, informer)

	// in-memory-cacheに追加？？
	indexer.Add(&v1.Pod{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:      "mypod",
			Namespace: v1.NamespaceDefault,
		},
	})

	// Now let's start the controller
	stop := make(chan struct{})
	defer close(stop)
	go controller.Run(1, stop)

	// Wait forever
	select {}
}
