package main

import "fmt"

// 関数を戻り値とする関数

func adder() func(int) int { // 引数なしの戻り値にint型を返すadder関数
	sum := 0
	return func(x int) int { //戻り値として無名関数（クロージャー）を返す
		sum += x // sum = sum + x
		return sum
	}
}

func main() {
	a := adder() // aは func (x int) int

	// fmt.Println(a(1))

	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
		/* ここのaに
		sum += x // sum = sum + x
		return sum
		の処理が入るイメージ
		*/
	}
}
