package main

import (
	"fmt"
	"reflect"
)

func main() {
	tolerations_a := []Toleration{
		{Key: "zone-y", Operator: "Exists"},
		{Key: "zone-z", Operator: "Exists"},
		{Key: "zone-common", Operator: "Exists"},
	}

	tolerations_b := []Toleration{
		{Key: "zone-y", Operator: "Exists"},
		{Key: "zone-z", Operator: "Exists"},
		{Key: "zone-common", Operator: "Exists"},
	}

	tolerations_c := []Toleration{
		{Key: "zone-y", Operator: "Exists"},
		{Key: "zone-z", Operator: "Exists"},
		{Key: "zone-q", Operator: "Exists"},
	}

	diff := reflect.DeepEqual(tolerations_a, tolerations_b)
	fmt.Println("diff:", diff)

	diff = reflect.DeepEqual(tolerations_a, tolerations_c)
	fmt.Println("diff:", diff)

}

// The pod this Toleration is attached to tolerates any taint that matches
// the triple <key,value,effect> using the matching operator <operator>.
type Toleration struct {
	Key string `json:"key,omitempty" protobuf:"bytes,1,opt,name=key"`

	Operator string `json:"operator,omitempty" protobuf:"bytes,2,opt,name=operator,casttype=TolerationOperator"`

	Value string `json:"value,omitempty" protobuf:"bytes,3,opt,name=value"`

	Effect string `json:"effect,omitempty" protobuf:"bytes,4,opt,name=effect,casttype=TaintEffect"`

	TolerationSeconds *int64 `json:"tolerationSeconds,omitempty" protobuf:"varint,5,opt,name=tolerationSeconds"`
}
