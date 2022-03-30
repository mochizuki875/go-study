/*
Goでキューを扱いたい
mainルーチンの中だけでやるのであればsliceとappendを使えばできる
*/

package main
import (
	"fmt"
)

func printSlice(s []int){
	fmt.Printf("[Debug] len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	// 長さ0のint型配列作成（これがキューに当たる）
	q := make([]int, 0)

	// キューに1, 2, 3を追加
	q = append(q, 1) // sliceへ新しい要素を追加
	printSlice(q)
	q = append(q, 2) // sliceへ新しい要素を追加
	printSlice(q)
	q = append(q, 3) // sliceへ新しい要素を追加
	printSlice(q)

	// 最初の値を取得
	v1 := q[0]
	q = q[1:] // 0番目の要素を切り捨てて再スライス
	fmt.Printf("q[0]=%d\n", v1)

	// 次の値を取得
	v2 := q[0]
	q = q[1:] // 0番目の要素を切り捨てて再スライス
	fmt.Printf("q[0]=%d\n", v2)

}