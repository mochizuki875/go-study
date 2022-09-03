package main
import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// 変数レシーバー
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// ポインタを引数に取る通常の関数
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
	v := Vertex{3,4}
	fmt.Println("v.Abs():", v.Abs())
	fmt.Println("AbsFunc(v):", AbsFunc(v))

	p := &Vertex{3,4}
	fmt.Println("p.Abs():", p.Abs()) //(*p).Absと解釈される
	fmt.Println("AbsFunc(*p):", AbsFunc(*p))
	
}