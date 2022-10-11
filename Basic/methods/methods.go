package main

import (
	"fmt"
	"math"
)

/* Goにクラスの概念はないが、
structに対してメソッドを紐付けることができる
*/

// 構造体の定義
type Vertex struct {
	X, Y float64
}

// メンバ関数のイメージ（methodと呼ぶ）
// ReceiverとしてVertex型のvを取る(引数なし)
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 通常の関数として実装するパターン
// 引数としてVertex型のvを取る
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println("v.Abs()=", v.Abs())       // メソッド呼び出し
	fmt.Println("AbsFunc(v)=", AbsFunc(v)) // 関数呼び出し

}
