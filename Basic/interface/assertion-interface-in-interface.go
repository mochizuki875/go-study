/*
構造体のフィールドにインターフェイスが含まれる場合の実験
*/

package main

import "fmt"

type IF_FOO interface {
	func_foo() string
}

type IF_BAR interface {
	func_bar() string
}

type IF_BAZ interface {
	func_foo() string
}

// Foo構造体メンバーにインターフェイスを含む
type Foo struct {
	BAR   IF_BAR
	BAZ   IF_BAZ
	value string
}

// IF_FOOの実装
func (f *Foo) func_foo() string {
	return "func foo"
}

// Bar構造体
type Bar struct {
	value string
}

// IF_BARの実装
func (f *Bar) func_bar() string {
	return "func bar"
}

// アサーション
func assertionFoo(i interface{}) {
	value, ok := i.(*Foo)
	fmt.Println(value)
	fmt.Println(ok)
	fmt.Println("------------------")
}

func assertionBar(i interface{}) {
	value, ok := i.(*Bar)
	fmt.Println(value)
	fmt.Println(ok)
	fmt.Println("------------------")
}

func assertionIFBar(i interface{}) {
	value, ok := i.(IF_BAR)
	fmt.Println(value)
	fmt.Println(ok)
	fmt.Println("------------------")
}

// 今回検証したいのはコレ
// iにIF_BAZで定義されているメソッドが実装されているかを確認する
func assertionIFBaz(i interface{}) {
	value, ok := i.(IF_BAZ)
	fmt.Println(value)
	fmt.Println(ok)
	fmt.Println("------------------")
}

func main() {
	var B IF_BAR
	bar := &Bar{value: "bar"}
	B = bar

	var F IF_FOO
	foo := &Foo{BAR: bar, value: "foo"} // Foo構造体初期化時のメンバーとしてBARを実装
	F = foo

	fmt.Println(B.func_bar())
	fmt.Println(F.func_foo())
	fmt.Println(foo.BAR.func_bar())

	// アサーション
	assertionFoo(B)       // false
	assertionFoo(F)       // true
	assertionBar(B)       // true
	assertionBar(foo.BAR) // true

	assertionIFBar(F)       // false
	assertionIFBar(B)       // true
	assertionIFBaz(F)       // true ★FにIF_BAZで定義されているメソッドが実装されているのでtrueになる
	assertionIFBar(foo.BAR) // true

}
