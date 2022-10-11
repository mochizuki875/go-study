package main

import (
	"fmt"
	"reflect"
)

type VertexA struct {
	X int
	Y int
}

type VertexB struct {
	X []int
	Y []int
}

// インターフェイス定義
type interfaceFunc interface {
	implFunc() int
}

// Implメソッドを定義
func (v *VertexA) implFunc() int {
	return v.X + v.Y
}

// Implメソッドを定義
func (v *VertexB) implFunc() int {
	sum := 0
	for _, value := range v.X {
		sum += value
	}
	for _, value := range v.Y {
		sum += value
	}
	return sum
}

func main() {
	var s interfaceFunc

	s1 := VertexA{3, 4}
	s2 := VertexB{[]int{1, 2, 3}, []int{4, 5, 6}}

	fmt.Println(s1)
	s = &s1 // 実装 ★実装はPointer-Receiverで定義されてるのでs=s1ではダメ
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(s.implFunc())

	fmt.Println("==================")

	fmt.Println(s2)
	s = &s2 // 実装
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(s.implFunc())

}
