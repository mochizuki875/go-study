package main

import "fmt"

// 定数の宣言
const j int = 42
const k = 69

// グループ化して定数の宣言
const (
	zero = iota // iota を初期値とすると 0 から始まる等差数列を作る
	one
	two
	three
	_
	five
)

func main(){
	fmt.Println(j,k)
	fmt.Println(zero,one,two,three,five)
}