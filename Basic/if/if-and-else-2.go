package main

import (
	"fmt"
)

func main() {

	type Data struct {
		a int
		b int
	}

	datas := []Data{
		{1, 3},
		{3, 1},
		{3, 3},
		{2, 2},
	}

	// aを3で割ったあまりが0またはbを3で割ったあまりが0のときXXXと表示する
	// さらにbを3で割ったあまりが0でない時はYYYと表示する

	// 既存パターン(冗長)

	fmt.Println("///////// 既存パターン /////////")
	for i, d := range datas {
		fmt.Printf("---- ケース%d: (%d,%d) ---- \n", i+1, d.a, d.b)
		if d.a%3 == 0 || d.b%3 == 0 {
			fmt.Println("XXX")
		}
		if d.b%3 != 0 {
			fmt.Println("YYY")
		}
	}

	fmt.Println("///////// 改善パターン /////////")
	for i, d := range datas {
		fmt.Printf("---- ケース%d: (%d,%d) ---- \n", i+1, d.a, d.b)

		if d.b%3 == 0 {
			fmt.Println("XXX")
		} else {
			if d.a%3 == 0 {
				fmt.Println("XXX")
			}
			fmt.Println("YYY")
		}
	}

}
