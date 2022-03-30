package main
import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ポインタレシーバーとすることでレシーバー自身を更新する
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 変数レシーバーの場合は値をコピーして操作（元の値は変わらない）
func (v Vertex) ScaleFunc(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3,4}
	fmt.Println("v:", v)
	v.Scale(10) // vの持つ値を10倍する
	fmt.Println("v:", v)
	fmt.Println("v.Abs():", v.Abs())

	// 変数レシーバーを持つ関数なので元の値は変わらない
	v.ScaleFunc(100) // vのを100倍する
	fmt.Println("v:", v)
	
}