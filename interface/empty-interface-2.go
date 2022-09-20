package main

import (
	"fmt"
	"reflect"
)

func Calc(obj interface{}) {

	fmt.Println("Type of arg: ", reflect.TypeOf(obj))
	fmt.Printf("(%v, %T)\n", obj, obj)

	// 型の変換とチェック
	// 引数をint型として扱う場合
	value, ok := obj.(int)
	fmt.Printf("(%v, %T)\n", value, value)
	fmt.Printf("(%v, %T)\n", ok, ok)
}

// 空インターフェイスはどんな型の値でも代入可能
func main() {
	Calc(1)
	fmt.Println("==============")
	Calc("Hello")
	fmt.Println("==============")
	fn := Calc
	fn(nil)
}
