package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string

	// 各型について初期化しなかった場合の値を出力
	fmt.Printf("Type: %T , ZeroValue: %v\n", i, i)
	fmt.Printf("Type: %T , ZeroValue: %v\n", f, f)
	fmt.Printf("Type: %T , ZeroValue: %v\n", b, b)
	fmt.Printf("Type: %T , ZeroValue: %v\n", s, s)

}
