/*
解説
2つの型（MyFloatとVertex）を用意して、型に応じて処理内容は違うが共通のメソッドを持たせたい
この時インターフェイス（Abser）という形で仮のメソッドを定義しておき、
その中で呼ばれる処理（Abs()）を各型ごとに実装することで、
実際の処理をインターフェイスから型に応じて呼び分けることができる
*/

package main

import (
	"fmt"
	"math"
	"reflect"
)

// MyFloat型の定義
type MyFloat float64

// Vertex型の定義
type Vertex struct {
	X, Y float64
}

// Abserインターフェイスの定義（Myfloat、Vertex型共通で使う）
type Abser interface {
	Abs() float64
}

// Myfloat型のインターフェイス実装
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// &Vertex型のインターフェイス実装
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser // ★インターフェースAbser型の変数を用意

	f := MyFloat(-100)
	v := Vertex{3, 4}

	a = f // MyFloat型のfにインターフェイスを実装
	fmt.Println("a = ", a)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println("MyFloatのa.Abs(): ", a.Abs()) // ★MyFloat型にとってのAbs()が呼ばれる

	fmt.Println("==================")

	a = &v // &Vertex型の&vにインターフェイスを実装(AbsはPointer-Receiverなのでvとしてはダメ)
	fmt.Println("a = ", a)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println("*Vertexのa.Abs(): ", a.Abs()) // ★Vertex型にとってのAbs()が呼ばれる

	// Vertex型に対するインターフェイス実装がされていないためエラーになる
	// a = v
	// fmt.Println("Vertexのa.Abs(): ", a.Abs())
}
