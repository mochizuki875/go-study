package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil { // 宣言しただけで初期化していないsliceはnilと判定される
		fmt.Println("nil slice")
	}
}
