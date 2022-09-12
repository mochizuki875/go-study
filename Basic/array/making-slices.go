package main

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%s, len=%d, cap=%d %v\n", s, len(x), cap(x), x)
}

func main() {
	// makeを使ってゼロ化された動的配列を作る
	// 動的配列を定義する場合はこれ
	a := make([]int, 5) // len=5の配列(slice)を作成
	fmt.Println("a = ", a)
	printSlice("a", a) // (length, capacity)=(5,5)

	fmt.Println("==========================")

	// 上の例と同義
	var x [5]int
	fmt.Println("x = ", x)
	fmt.Println("len=", len(x), "cap=", cap(x), x) // (length, capacity)=(5,5)

	fmt.Println("==========================")

	b := make([]int, 0, 5) // len=0, cap=5の配列
	fmt.Println("b = ", b)
	printSlice("b", b) // (length, capacity)=(0,5)

	fmt.Println("==========================")

	c := b[:2] // bからslice、bはlen=0だがcap=5なので拡張してslice
	fmt.Println("c = ", c)
	printSlice("c := b[:2]", c) // (length, capacity)=(2,5)

	fmt.Println("==========================")

	d := b[:cap(b)] // bからslice、bはlen=0だがcap=5なので拡張してslice
	fmt.Println("d = ", d)
	printSlice("e := b[:cap(b)]", d) // (length, capacity)=(5,5)

	fmt.Println("==========================")

	// len=0に対してこれはダメっぽい
	// e := b[2:] // bからslice
	// fmt.Println("e = ", e)
	// printSlice("e := b[2:]", e)

	fmt.Println("==========================")

	f := b[2:5] // bからslice
	fmt.Println("f = ", f)
	printSlice("f := b[2:5]", f) // (length, capacity)=(3,3)

}
