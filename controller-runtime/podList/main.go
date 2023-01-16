package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"log"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func main() {

	scheme := runtime.NewScheme()
	ctx := context.Background()

	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	log.Println("Create Manager.")
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     ":8080",
		Port:                   9443,
		HealthProbeBindAddress: ":8081",
		LeaderElection:         false,
		LeaderElectionID:       "0c174af9.example.com",
	})
	if err != nil {
		log.Println("Create Manager faild.: ", err)
		os.Exit(1)
	}
	log.Println("Manager Created.")

	log.Println("Start Manager goroutine.")
	go mgr.Start(ctrl.SetupSignalHandler())

	log.Println("Wait for Manager setup 5s.")
	time.Sleep(time.Second * 5)

	log.Println("Get client.")
	k8sclient := mgr.GetClient()

	log.Println("Get PodList using client.")
	var podList corev1.PodList
	opt := client.ListOptions{Namespace: "default"}

	if err = k8sclient.List(ctx, &podList, &opt); err != nil {
		log.Println("List error.: ", err)
	}

	// show pod object to stdout
	for i, pod := range podList.Items {
		fmt.Printf("[Pod Name %d]%s\n", i, pod.GetName())
	}

}
