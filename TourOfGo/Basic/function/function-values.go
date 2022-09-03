package main

import (
	"fmt"
)

// 関数を引数とする関数

// 引数で指定された関数に所定の値を入れた際の戻り値を返す
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	// 関数の定義
	hypot := func(x, y float64) float64 {
		return x + y
	}

	// hypot関数の動作を確認
	fmt.Println(hypot(3, 4))

	// hypot関数をcompute関数に渡した時の動作
	fmt.Println(compute(hypot))
}
