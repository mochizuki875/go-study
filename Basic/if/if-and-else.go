package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { // powはべき乗
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim // limを超えたらlimを返す
}

func main() {
	fmt.Println(pow(3, 2, 10))
	fmt.Println("============")
	fmt.Println(pow(3, 3, 20))
}
