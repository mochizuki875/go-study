package main

import "fmt"

func main() {

	// arrayを定義
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println("names= ", names)

	// sliceを取得
	slice_a := names[0:2]
	slice_b := names[1:3]

	fmt.Println("slice_a= ", slice_a)
	fmt.Println("slice_b= ", slice_b)

	slice_a[1] = "XXX"

	// sliceは配列の指定した要素を参照しているだけ
	// 値を書き換えると配列自体の値が更新される
	fmt.Println("names= ", names)
	fmt.Println("slice_a= ", slice_a)
	fmt.Println("slice_b= ", slice_b)

}
