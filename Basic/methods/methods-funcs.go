package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func (v Vertex) add() int {
	return v.X + v.Y
}

func (v Vertex) update() (int, int) {
	return v.X * 2, v.Y * 2
}

func ADD(v Vertex) int {
	return v.X + v.Y
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.add())
	fmt.Println(ADD(v))

	fmt.Println(v.update()) // 6 8
	fmt.Println(v)          // {3, 4} 元の構造体の値は変わらない
}
