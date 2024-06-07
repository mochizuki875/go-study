package main

import "fmt"

type str_x struct {
	a int
	b int
}

func (s str_x) sum() int {
	return s.a + s.b
}

type str_y struct {
	str_x
	c int
}

func main() {
	var sum int

	x := str_x{a: 1, b: 2}
	y := str_y{str_x: x, c: 3}

	// sum()を呼び出す時ってどっちの書き方が一般的？？
	sum = y.sum() // パターン①
	fmt.Println("Sum: ", sum)
	sum = y.str_x.sum() // パターン②（str_xとstr_yで同名のメソッドを持つ時は明示的にstr_xを指定する。普通やらない）
	fmt.Println("Sum: ", sum)

}
