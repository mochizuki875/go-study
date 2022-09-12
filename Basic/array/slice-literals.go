package main

import "fmt"

func main() {

	// 配列のリテラル(初期化) → 要素数を指定
	a := [6]int{0, 1, 2, 3, 4, 5}
	fmt.Println("a= ", a)

	// sliceのリテラル(初期化) → 要素数を指定しない
	// q := []int{2, 3, 5, 7, 11, 13}
	var q []int = []int{2, 3, 5, 7, 11, 13}
	fmt.Println("q= ", q)

	// sliceのリテラル(初期化) → 要素数を指定しない
	r := []bool{true, false, true, true, false, true}
	fmt.Println("r= ", r)

	// i,bをフィールドに持つ構造体を要素にもつslice
	s := []struct {
		i int
		b bool
	}{
		{2, true}, // sliceを要素として取得
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println("s= ", s)

	// 上記は以下と同義
	// 構造体を要素に持つslice
	// https://nishinatoshiharu.com/initialize-structure-slice/

	// type V struct {
	// 	i int
	// 	b bool
	// }

	// var s []V = []V{
	// 	{2, true}, // sliceを要素として取得
	// 	{3, false},
	// 	{5, true},
	// 	{7, true},
	// 	{11, false},
	// 	{13, true},
	// }

	// fmt.Println("s= ", s)
}
