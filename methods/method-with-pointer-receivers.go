package main
import (
	"fmt"
	"math"	
)

type Vertex struct {
	X,Y float64
}

// ポインタレシーバ
func (v *Vertex) Scale(f float64) {
	v.X = v.X + f
	v.Y = v.Y + f
}

// ポインタレシーバ
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 値レシーバ
func (v Vertex) Mult(f float64) Vertex {
	v.X = v.X * f
	v.Y = v.Y * f

	var a Vertex
	a.X = v.X
	a.Y = v.Y
	return a
}

func main() {
	v := &Vertex{3,4}
	
	fmt.Printf("Before Scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After Scaling: %+v, Abs: %v\n", v, v.Abs())

	v1 := Vertex{1,3}
	fmt.Printf("Before Scaling: %+v, Abs: %v\n", v1, v1.Abs())
	v1.Scale(5)
	fmt.Printf("After Scaling: %+v, Abs: %v\n", v1, v1.Abs())

	// Multは値レシーバなのでv1の値は直接更新されない
	v2 := v1.Mult(10)
	fmt.Printf("v1: %+v, v2 Mult: %+v\n", v1, v2)

}