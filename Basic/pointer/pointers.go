/*
	&はアドレスを示す
	*はポインタの示す先の値を示す
	var p *int はポインタ変数を示す(アドレスを格納)

*/

package main

import "fmt"

func main() {
	i, j := 42, 2701

	var p *int // ポインタ変数p
	p = &i     // pにアドレスを格納
	// p := &i                  // (上記をまとめるパターン)point to i (*i型) &はメモリアドレスを示す
	fmt.Println("*p = ", *p) // *pはポインタpの指す値を示す
	fmt.Println("p = ", p)   // pはポインタの値（メモリアドレス）そのものを示す

	*p = 21 // pの示す先（i）に値を格納
	fmt.Println("i(*p) = ", i)

	p = &j       // pの示す先をjのアドレスに変更
	*p = *p / 37 // *pの値を演算
	fmt.Println("j = ", j)
	fmt.Println("&j = ", &j)
	fmt.Println("p = ", p)
}
