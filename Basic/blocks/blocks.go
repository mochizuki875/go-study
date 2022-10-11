/*
Blocks
実行スコープを分離する
https://golang.org/ref/spec#Blocks
https://play.golang.org/p/4wonhoSSqZi
*/

package main

import "fmt"

func main() {
	a := 1
	fmt.Println(a)
	a = a + 1
	fmt.Println(a)

	// 中括弧で括った部分はスコープが分かれる
	{
		a := 10
		fmt.Println(a)
		a = a + 10
		fmt.Println(a)

	}

}
