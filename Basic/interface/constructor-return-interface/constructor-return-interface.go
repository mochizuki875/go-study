/*
コンストラクタとしてinterfaceを返すパターンとポインターを返すパターンの比較
*/

package main

import (
	"fmt"
	"reflect"
)

type Values struct {
	X int
	Y int
}

// interface定義
type Interface interface {
	GetValues() Values
	SetValues(x, y int)
}

// メソッドの実装(Getterに相当)
func (v *Values) GetValues() Values {
	return *v
}

// メソッドの実装(Setterに相当)
func (v *Values) SetValues(x, y int) {
	v.X = x
	v.Y = y
}

// ポインターを返すコンストラクタ(的な)関数
func NewValuesPointer(x, y int) *Values {
	v := &Values{x, y}
	return v
}

// ★interfaceを返すコンストラクタ(的な)関数
// 戻り値としてInterfaceを、returnとして&Valuesを指定することで
// i I := t T[]をしていることになる
func NewValuesInterface(x, y int) Interface {
	v := &Values{x, y}
	return v
}

func main() {
	fmt.Println("===普通にinterfaceを使うパターン===")
	var i1 Interface
	v := &Values{2, 3}
	i1 = v
	fmt.Println(reflect.TypeOf(i1))
	// fmt.Println(*i1) これは不可
	fmt.Println("fmt.Println(*i1)みたいな直接参照は不可(i1はinterfaceなのでstruct自身は隠蔽されてる→GetterやSetterメソッドを使う)")
	fmt.Println("i1.GetValues() = ", i1.GetValues()) // Getterメソッドを使った参照

	fmt.Println("")

	fmt.Println("===interfaceを返すコンストラクタを使うパターン===")
	i2 := NewValuesInterface(2, 3)
	fmt.Println(reflect.TypeOf(i2))
	// fmt.Println(*i2) これは不可
	fmt.Println("fmt.Println(*i2)みたいな直接参照は不可(i1はinterfaceなのでstruct自身は隠蔽されてる→GetterやSetterメソッドを使う)")
	fmt.Println("i2.GetValues() = ", i2.GetValues()) // Getterメソッドを使った参照
	fmt.Println("i2.SetValues(4,6)")
	i2.SetValues(4, 6) // Setterメソッドを使って値を更新
	fmt.Println("i2.GetValues() = ", i2.GetValues())

	fmt.Println("")

	fmt.Println("===ポインタを返すコンストラクタを使うパターン===")
	i3 := NewValuesPointer(2, 3) // 直接参照
	fmt.Println(reflect.TypeOf(i3))
	fmt.Println("*i3 = ", *i3, " ※i3は構造体へのポインターなので直接参照可能")
	fmt.Println("i3.GetValues() = ", i3.GetValues()) // Getterメソッドを使った参照
	i3.X = 4                                         // 直接更新
	i3.Y = 6
	fmt.Println("*i3 = ", *i3)

	fmt.Println("")
}
