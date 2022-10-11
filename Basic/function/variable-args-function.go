/*
引数の数が可変な関数
https://mebee.info/2021/06/13/post-23392/
*/

package main

import "fmt"

// int型の引数を好きなだけ取れる
func sum(num ...int) int {
	result := 0
	for _, v := range num {
		result += v
	}
	return result
}

func main() {
	// 引数を直接指定するパターン
	fmt.Println(sum(1, 2, 3, 4, 5, 6))

	// 引数をsliceで指定するパターン
	num := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sum(num...))
}
