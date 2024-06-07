package main

import "fmt"

// struct定義
// structはフィールドの集合
type Vertex struct {
	X int
	Y int

	// The total number of nodes that should be running the daemon
	// pod (including nodes correctly running the daemon pod).
	// More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
	DesiredNumberScheduled int32 `json:"desiredNumberScheduled" protobuf:"varint,3,opt,name=desiredNumberScheduled"`

	// The number of nodes that should be running the
	// daemon pod and have one or more of the daemon pod running and
	// available (ready for at least spec.minReadySeconds)
	// +optional
	NumberAvailable int32 `json:"numberAvailable,omitempty" protobuf:"varint,7,opt,name=numberAvailable"`
}

func main() {
	v := Vertex{} // structの初期化
	v.X = 1
	v.Y = 2

	fmt.Println(v)
	// z := v.Z
	// fmt.Println(z)

	numberAvailable := v.NumberAvailable
	fmt.Println(numberAvailable)

}
