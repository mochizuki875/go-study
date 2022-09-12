package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Printf("(%v, %T)\n", s, s)

	s, ok := i.(string) // 型がstringであるかを確認(アサーション)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Printf("(%v, %T)\n", f, f)
	fmt.Println(f, ok)
}
