/*
パッケージを読み込んで使う例
fooというパッケージを定義して使っている
*/

package main

import (
	"fmt"

	"example.com/packagetest/foo"
)

func main() {
	f := foo.Foo{}
	f.A = 1
	f.B = 2
	// f.C = 3 // 不可
	// f.D = 4 // 不可
	f.Set_c(3)
	f.Set_d(4)
	fmt.Println("f: ", f)

}
