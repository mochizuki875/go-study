/*
非公開フィールド/メソッドの検証
先頭が小文字で始まるフィールドやメソッドは非公開(unexported)と呼ばれpackage外からアクセスできない
*/

package main

import (
	"fmt"
	"reflect"
	"unsafe"

	"example.com/unexportedfield/foo"
)

func main() {
	f := foo.Foo{}
	f.A = 1
	f.B = 2
	// f.C = 3 // 不可
	// f.D = 4 // 不可
	f.Set_c(3)
	f.Set_d(4)
	// c := get_c() // 不可
	fmt.Println("f: ", f)

	// 非公開フィールドの値を参照/更新

	var outer *foo.Foo // ポインタ型でstructを宣言
	outer = &foo.Foo{}
	// outer := &foo.Foo{}
	outer.A = 1
	outer.B = 2
	outer.Set_c(3)
	outer.Set_d(4)

	fmt.Printf("outer: %#v\n", *outer) // 初期値表示

	var ref_outer reflect.Value = reflect.ValueOf(outer).Elem()
	fmt.Printf("ref_outer: %#v\n", ref_outer) // 複製した値を表示

	var ref_c reflect.Value = ref_outer.FieldByName("c")
	c := (*int)(unsafe.Pointer(ref_c.UnsafeAddr())) // 非公開フィールドのポインタを取得
	fmt.Printf("c: %#v\n", *c)                      // 取得した非公開フィールドの値を表示
	*c = 30                                         // 非公開フィールドを更新
	fmt.Printf("c: %#v\n", *c)                      // 更新後の非公開フィールドの値を表示

	fmt.Printf("outer: %#v\n", *outer) // 更新後の値を表示

}
