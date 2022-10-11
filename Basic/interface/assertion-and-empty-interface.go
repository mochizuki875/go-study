/*
  Go言語には、全ての型と互換性を持っているinterface{}型（空インターフェース）がある
  また、アサーションによりインターフェイスを実装した実態の型を取得・確認することができる。
  https://blog.y-yuki.net/entry/2017/05/08/000000
*/

package main

import (
	"fmt"
	"reflect"
)

// 引数の型をinterface{}にすると、どんな型の値でも受け取ることができる関数を記述可能
func Calc(obj interface{}) {

	// interface{}型の引数で受け渡された値は元の型の情報が欠落する
	fmt.Println("Type of arg: ", reflect.TypeOf(obj))
	fmt.Printf("obj(value, type) = (%v, %T)\n", obj, obj)

	// アサーション
	// アサーションを用いることで実態の型への変換とチェックを行える
	// 引数をint型として扱う場合
	value, ok := obj.(int)
	fmt.Printf("value(value, type) = (%v, %T)\n", value, value)
	fmt.Printf("ok(value, type) = (%v, %T)\n", ok, ok)
}

func main() {
	// 空インターフェイスはどんな型の値でも代入可能
	var i interface{} // 空のインターフェイス
	fmt.Printf("(%v, %T)\n", i, i)
	fmt.Println("==============")

	// int32を代入
	i = 42
	fmt.Printf("(%v, %T)\n", i, i)
	fmt.Println("==============")

	// stringを代入
	i = "hello"
	fmt.Printf("(%v, %T)\n", i, i)
	fmt.Println("==============")

	fmt.Println("=== 空インターフェイスを引数に取る関数 ===")

	// 空インターフェイスを引数に取る関数にはどんな型も代入可能
	Calc(1)
	fmt.Println("==============")
	Calc("Hello")
	fmt.Println("==============")
	fn := Calc
	fn(nil)

}
