/*
解説
interfaceにnilも実装できるという例
*/

package main

import (
	"fmt"
	"reflect"
)

// Vertex型の定義
// type Vertex struct {
// 	X, Y float64
// }

// Abserインターフェイスの定義（Myfloat、Vertex型共通で使う）
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser // ★インターフェースAbser型の変数を用意

	a = nil
	fmt.Println("a = ", a)
	fmt.Println(reflect.TypeOf(a))

}
