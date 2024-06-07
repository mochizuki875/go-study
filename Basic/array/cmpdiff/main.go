package main

import (
	"fmt"
	"strings"

	"github.com/google/go-cmp/cmp"
)

func main() {
	// 比較元
	tolerations_a := []Toleration{
		{Key: "zone-y", Operator: "Exists"},
		{Key: "zone-z", Operator: "Exists"},
		{Key: "zone-common", Operator: "Exists"},
	}

	// 比較対象
	tolerations_b := []Toleration{
		{Key: "zone-y", Operator: "Exists"},
		{Key: "zone-z", Operator: "Exists"},
		{Key: "zone-common", Operator: "Exists"},
		{Key: "xxx", Operator: "Exists"},
		{Key: "yyy", Operator: "Exists"},
	}

	fmt.Println("tolerations_b:", tolerations_b)

	var prefixedTolerations_b []Toleration

	for _, toleration := range tolerations_b {
		if strings.HasPrefix(toleration.Key, "zone-") {
			prefixedTolerations_b = append(prefixedTolerations_b, toleration)
		}
	}

	fmt.Println("prefixedTolerations_b:", prefixedTolerations_b)

	diff := cmp.Diff(tolerations_a, prefixedTolerations_b)
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
