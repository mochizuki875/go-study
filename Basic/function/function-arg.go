/*
関数を引数に取るパターン
*/

package main

import "fmt"

type Fn func([]int) // 関数をtypeとして定義

func Add(fn Fn) { // 関数型を引数に取る関数
	values := []int{1, 2, 3, 4, 5}
	fmt.Println(values)
	fn(values)
}

func main() {

	Add(func(values []int) { // 引数として関数を定義
		sum := 0
		for _, v := range values {
			sum += v
		}
		fmt.Println(sum)
	})

}
