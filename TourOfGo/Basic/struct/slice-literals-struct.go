package main

import "fmt"

func main() {

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
