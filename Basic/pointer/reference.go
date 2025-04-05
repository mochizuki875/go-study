package main

import "fmt"

type str struct {
	x int
	y int
	c string
}

func main() {
	var s *str
	s = &str{10, 20, "a"}
	fmt.Printf("s: %+v\n", s)
	mod(s)

	fmt.Printf("s: %+v\n", s)

}

func mod(str *str) {

	// copy := *str // 直接strのメモリを書き換えないためにコピーを作成
	// copy := str // この場合はstrのメモリをcopyが参照するため、copyを書き換えるとstrも書き換わる

	copyObj := *str
	copy := &copyObj
	// copy := &(*str) // これをやると直接参照になる

	copy.x = 100
	copy.y = 200
	copy.c = "z"

	fmt.Printf("copy: %+v\n", *copy)
}
