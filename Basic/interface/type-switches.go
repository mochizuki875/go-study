package main

import "fmt"

// 空インターフェイスを引数にするとどんな値でも引数に取れる
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v)
		fmt.Printf("int, (%v, %T)\n", v, v)
	case string:
		fmt.Printf("string, (%v, %T)\n", v, v)
	default:
		fmt.Printf("other, (%v, %T)\n", v, v)
	}
}

func main() {
	do(1)
	do("hello")
	do(true)
}
