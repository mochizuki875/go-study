package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { // if文の条件の前に変数宣言を含めることができる
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func calc(max int) int {
	var sum = 1
	for sum < max {
		if sum += sum; sum >= max {
			break
		} else {
			fmt.Println("sum = ", sum)
		}
	}
	return sum
}

func main() {
	// fmt.Println(
	// 	pow(3,2,10),
	// 	pow(3,3,20),
	// )
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 10))

	fmt.Println("calcResult = ", calc(100))
}
