package main

import "fmt"

// struct定義
// structはフィールドの集合
type Vertex struct {
	X int
	Y int
}

func main() {
	// fmt.Println(Vertex{1,2})
	v := Vertex{1, 2} // structの初期化
	v.X = 4           // structのフィールドにアクセス
	fmt.Println(v.X)
	fmt.Println(v)
}
