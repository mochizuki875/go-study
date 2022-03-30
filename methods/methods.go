package main
import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// メンバ関数のイメージ（methodと呼ぶ）
// 引数としてVertex型のvを取る（receiver）
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 通常の関数として実装するパターン
// 引数としてVertex型のvを取る
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main(){
	v := Vertex{3,4}
	fmt.Println("v.Abs()=", v.Abs()) // 引数の指定なし
	fmt.Println("AbsFunc(v)=", AbsFunc(v)) // 引数としてVertex型を渡す
	
}