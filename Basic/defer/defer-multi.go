package main

import "fmt"

func main() {
	fmt.Println("counting") // 1番目に実行

	for i := 0; i < 10; i++ {

		// 最後に実行
		// LIFO(最後に呼び出されたものから)の順番で実行
		defer fmt.Println(i)
	}
	fmt.Println("done") // 2番目に実行
}
