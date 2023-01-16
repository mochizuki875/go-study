package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"log"

	"github.com/google/go-cmp/cmp"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func main() {

	scheme := runtime.NewScheme()
	ctx := context.Background()

	podName := "test-pod"
	podNamespace := "default"

	// 標準リソースの型をschemeに追加
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	// Create ControllerManager
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
		log.Fatalln("Faild to Create Manager.: ", err)
		os.Exit(1)
	}
	log.Println("Manager Created.")

	// Start ControllerManager
	log.Println("Start Manager.")
	go mgr.Start(ctrl.SetupSignalHandler())

	time.Sleep(time.Second * 5)

	// Create Client
	// client.Clientインターフェイスを取得
	// https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/client#Client
	log.Println("Get client.")
	k8sClient := mgr.GetClient()

	// Delete old Pod
	log.Println("Delete old Pod.")
	u := &unstructured.Unstructured{}
	u.SetName(podName)
	u.SetNamespace(podNamespace)
	u.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "",
		Kind:    "Pod",
		Version: "v1",
	})
	if err = k8sClient.Delete(ctx, u); client.IgnoreNotFound(err) != nil {
		log.Fatalln("Faild to Delete() Pod.: ", err)
		os.Exit(1)
	}
	time.Sleep(time.Second * 5)

	// Create Pod
	log.Println("Create test pod.")
	podTemplate := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podName,
			Namespace: podNamespace,
			Annotations: map[string]string{
				"app":       "test-Pod",
				"podStatus": "Created",
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "nginx",
					Image: "nginx:latest",
				},
			},
		},
	}

	if err = k8sClient.Create(ctx, podTemplate); err != nil {
		log.Fatalln("Faild to Create() Pod.: ", err)
		os.Exit(1)
	}

	time.Sleep(time.Second * 5)

	// Get Pod
	var basePodInst, podInst corev1.Pod

	if err = k8sClient.Get(ctx, client.ObjectKey{
		Namespace: podNamespace,
		Name:      podName,
	}, &basePodInst); err != nil {
		log.Fatalln("Faild to Get() Pod.: ", err)
		os.Exit(1)
	}
	fmt.Println(basePodInst.GetName())

	// DeepCopy
	podInst = *basePodInst.DeepCopy()

	// Diff between base and patched instance
	fmt.Println("=== [before Update] Show diff before Patch === ")
	log.Println(cmp.Diff(basePodInst, podInst))
	fmt.Println("============================== ")

	// Update
	podInst.Annotations["podStatus"] = "Patched"

	fmt.Println("=== [After Update] Show diff after Patch === ")
	log.Println(cmp.Diff(basePodInst, podInst))
	fmt.Println("============================== ")

	// Patch Pod
	// 第３引数はPatchを適用するためのPatchインターフェイス
	// https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/client#Patch
	if err = k8sClient.Patch(ctx, &podInst, client.MergeFrom(&basePodInst)); err != nil {
		log.Fatalln("Faild to Patch() Pod.: ", err)
		os.Exit(1)
	}
}
