package main

import "fmt"

// 関数名（引数）戻り値
func add(x int, y int) int {
	return x + y
}

func delta(x, y float64) float64 {
	return x - y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(delta(42.5, 13))
}
