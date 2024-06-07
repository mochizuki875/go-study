package main

import "fmt"

func main() {
	primes := [6]int{0, 1, 2, 3, 4, 5} // arrayは固定長
	// var primes2 []int

	// sliceは可変長でarrayを切り取る
	// []T := a[low : high]
	// low以上、high未満の要素を取得
	// arrayの要素は0スタート

	var s0 []int = primes[:7] // 7要素目まで
	fmt.Println(s0)

	fmt.Println(len(primes2))
	var s0 []int = primes2[:len(primes2)]
	fmt.Println(s0)

	// var a int32
	// var b int32

	// a = 5
	// b = 1

	// c := int(a - b)

	// fmt.Println(a, b, c)
}
