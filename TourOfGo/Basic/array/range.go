package main

import (
	"fmt"
)

func main() {

	var s = []int{0, 10, 20, 30, 40, 50, 60, 70, 80}
	fmt.Println(s)

	// rangeを用いることで要素数分forを回せる
	for i, v := range s {
		fmt.Printf("i = %d , v = %d\n", i, v)
	}
}
