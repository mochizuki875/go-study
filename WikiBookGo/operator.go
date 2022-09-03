package main

import "fmt"

func main() {
	a, b, c := 42, 12, 3
	t, f := true, false

	/* 単項演算子 */
	fmt.Println("単項演算子")
	fmt.Println(+c)
	fmt.Println(-c) // 単項マイナス
	fmt.Println(^c) // ビット反転（C言語系の ~ ではなく ^）
	fmt.Println(!t) // 論理否定
	fmt.Println("===================")

	/* 二項演算子 */
	fmt.Println("二項演算子")
	fmt.Println(a * b)  // 乗算
	fmt.Println(a / b)  // 除算
	fmt.Println(a % b)  // 剰余
	fmt.Println(a << c) // 左シフト
	fmt.Println(a >> c) // 右シフト
	fmt.Println(a & c)  // ビットごとの論理積
	fmt.Println(a &^ c) // ビットクリア（ &^ でトークン）
	fmt.Println(a & ^c) // ビットクリア（これとおなじ）
	fmt.Println("===================")

	fmt.Println(a + c) // 加算
	fmt.Println(a - c) // 減算
	fmt.Println(a | c) // ビットごとの論理和
	fmt.Println(a ^ c) // ビットごとの排他的論理和
	fmt.Println("===================")

	fmt.Println(a == b) // 一致
	fmt.Println(a != b) // 不一致
	fmt.Println(a < b)  // より小
	fmt.Println(a <= b) // より小または一致
	fmt.Println(a > b)  // より大
	fmt.Println(a >= b) // より大または一致
	fmt.Println("===================")

	fmt.Println(t && f) // 論理積（短絡評価あり）

	fmt.Println(t || f) // 論理和（短絡評価あり）
}
