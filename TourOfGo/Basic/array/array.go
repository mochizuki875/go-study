package main

import "fmt"

func main(){
	var a [2]string // 配列は固定長
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// primes := [6]int{6,5,4,3,2,1}
	var primes [6]int // 配列は固定長

	for i, _ := range primes { // rangeでは{index, element}が返る
		primes[i] = i
	}
	fmt.Println(primes)

}