package main

import (
	"fmt"
)

func main() {
	s := make([]int, 10) // length=10のslice
	// var s [10]int
	for i := range s { // 第2戻り値は省略可能
		s[i] = 1 << uint(i) // == 2**i
	}

	for _, value := range s {
		fmt.Printf("%d\n", value)
	}
}
