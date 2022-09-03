package main

import "fmt"

func main() {
	primes := [6]int{0, 1, 2, 3, 4, 5} // arrayは固定長

	// sliceは可変長でarrayを切り取る
	// []T := a[low : high]
	// low以上、high未満の要素を取得
	// arrayの要素は0スタート

	var s0 []int = primes[0:4] // 0要素目~3要素目まで
	fmt.Println(s0)

	var s1 []int = primes[1:4] // 1要素目~3要素目まで
	fmt.Println(s1)

	var s3 []int = primes[2:3] // 2要素目~2要素目まで
	fmt.Println(s3)

	var s4 []int = primes[2:2] // 何も取得しない
	fmt.Println(s4)
}
