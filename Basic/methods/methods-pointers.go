package main

import (
	"fmt"
)

// 構造体の定義
type Vertex struct {
	X, Y float64
}

// Pointer Receiver(こっちの方が一般的)
// 一般的にメソッドを実行して値を更新(直接参照)したいケースが多い
// Value Receiverだと値のコピーが発生する(間接参照)
// 構造体のポインタをReceiverとすることで構造体自身を更新する
func (v *Vertex) Add() {
	// v.X = v.X + v.X
	v.X += v.X
	// v.Y = v.Y * v.Y
	v.Y += v.Y
}

// Value Receiver
// 変数レシーバーの場合は値をコピーして操作（元の値は変わらない）
func (v Vertex) AddFunc() {
	v.X += v.X
	v.Y += v.Y
}

func main() {
	v1 := Vertex{3, 4}
	fmt.Println("v1:", v1)
	// ポインタReceiverのケース
	v1.Add()
	fmt.Println("v1:", v1) // メソッド実行によりv1の値が更新される

	fmt.Println("=============")

	v2 := Vertex{3, 4}
	fmt.Println("v2:", v2)
	// 変数レシーバーを持つ関数なので元の値は変わらない
	v2.AddFunc()
	fmt.Println("v2:", v2) // メソッド実行前後でv2の値に変化なし

}
