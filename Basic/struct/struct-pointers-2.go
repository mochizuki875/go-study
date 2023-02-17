package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v // structへのポインタ定義

	(*p).X = 50 // ポインタの示すフィールドにアクセス（本来はこう書く）
	fmt.Println(v)

	fmt.Println("*p = ", *p)
	fmt.Println("p = ", p)

	p.X = 100 // こう書いてもフィールドアクセス可能
	fmt.Println(v)
}
