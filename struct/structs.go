package main

import "fmt"

// struct定義
type Vertex struct {
	X int
	Y int
}

func main() {
	// fmt.Println(Vertex{1,2})
	v := Vertex{1, 2} // 構造体の初期化
	v.X = 4 // 構造体のフィールドにアクセス
	fmt.Println(v.X)
	fmt.Println(v)
}
