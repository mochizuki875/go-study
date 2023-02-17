package main

import "fmt"

type Foo struct {
	A int
	B int
}

func main() {
	var foo *Foo // ポインタ型の変数として構造体を宣言
	foo = &Foo{} // 初期化
	// foo = &Foo{1,2} // これでも良い

	foo.A = 1
	foo.B = 2

	fmt.Println("foo: ",foo)
	fmt.Println("*foo: ",*foo)
}