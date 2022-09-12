package main

import (
	"fmt"
)

// 任意の型を定義
type MyFloat float64

// MyFloat型に対してメソッドを定義
func (f MyFloat) Abs() float64 { // 絶対値を返す
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-100)
	var f2 MyFloat = 200

	fmt.Println("f:", f)
	fmt.Println("f.Abs():", f.Abs())

	fmt.Println("f2:", f2)
	fmt.Println("f2.Abs():", f2.Abs())
}
