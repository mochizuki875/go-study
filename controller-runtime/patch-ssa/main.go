package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"log"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	applycorev1 "k8s.io/client-go/applyconfigurations/core/v1"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/utils/pointer"
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
	var podInst corev1.Pod

	if err = k8sClient.Get(ctx, client.ObjectKey{
		Namespace: podNamespace,
		Name:      podName,
	}, &podInst); err != nil {
		log.Fatalln("Faild to Get() Pod.: ", err)
		os.Exit(1)
	}
	fmt.Println(podInst.GetName())

	// ApplyConfiguration型で変更を作成
	patchInst := applycorev1.Pod(podName, podNamespace).WithAnnotations(map[string]string{"podStatus": "ssa"})
	obj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(patchInst)
	if err != nil {
		log.Fatalln("Faild to convert unstructured.: ", err)
		os.Exit(1)
	}

	// Unstructuredに変換
	patch := &unstructured.Unstructured{
		Object: obj,
	}

	// 差分比較(今回の場合は最初にPod新規作成して必ず差分発生するのでなくて良い)
	// ApplyConfigurationに変換
	currentApplyConfig, err := applycorev1.ExtractPod(&podInst, "sample-client")
	if err != nil {
		log.Fatalln("Faild to create currentApplyConfig.: ", err)
		os.Exit(1)
	}
	if equality.Semantic.DeepEqual(patchInst, currentApplyConfig) {
		log.Println("Non diff.")
		os.Exit(0)
	}

	// Server-Side Apply
	// client-sampleとしてAnnotationを更新
	// managedFieldsとしてmanager=client-sampleで登録される
	err = k8sClient.Patch(ctx, patch, client.Apply, &client.PatchOptions{
		FieldManager: "client-sample",
		Force:        pointer.Bool(true),
	})
	if err != nil {
		log.Fatalln("Faild to Patch() Pod.: ", err)
		os.Exit(1)
	}

}
