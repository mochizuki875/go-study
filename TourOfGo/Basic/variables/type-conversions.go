package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y)) // float64に変換
	var z uint = uint(f)
	// var f float64 = math.Sqrt(x*x + y*y)
	// var z uint = f

	fmt.Println(x, y, f, z)
}
