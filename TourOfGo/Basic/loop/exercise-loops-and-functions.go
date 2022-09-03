package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)

	for z*z <= x {		
		z++
	}
	return z-1
}


func main() {
	fmt.Println("Ans: ", Sqrt(4.1))
}