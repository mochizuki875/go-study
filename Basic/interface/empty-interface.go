package main

import "fmt"

// 空インターフェイスはどんな型の値でも代入可能
func main() {
	var i interface{} // 空のインターフェイス
	fmt.Printf("(%v, %T)\n", i, i)

	i = 42
	fmt.Printf("(%v, %T)\n", i, i)

	i = "hello"
	fmt.Printf("(%v, %T)\n", i, i)
}
