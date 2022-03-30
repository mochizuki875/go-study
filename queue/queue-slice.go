/*
Goでキューを扱いたい
mainルーチンの中だけでやるのであればsliceとappendを使えばできる
キューをtypeで定義し、キューに対する処理をメソッドとして実装
*/

package main
import (
	"fmt"
)

type queue []int

func CreateQueue() queue {
	// q := make([]int, 0)
	var q queue
	return q
}

// キューに値を追加するメソッド
func (q *queue) EnQueue(t int) {
	fmt.Printf("q = append(q,%d)\n", t)
	*q = append(*q,t) // sliceへ新しい要素を追加
	printSlice(*q)
}

// キューから値を取り出すメソッド
func (q *queue) DeQueue() int {
	t := (*q)[0] // sliceの先頭の値を取り出す（FIFO）
	*q = (*q)[1:] // 0番目の要素を切り捨てて再スライス
	printSlice(*q)
	return t
}

// [Debug] slice情報確認用関数
func printSlice(s []int){
	fmt.Printf("[Debug] len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	q := CreateQueue()

	fmt.Println("EeQueue: ", 1)	
	q.EnQueue(1)
	fmt.Println("EeQueue: ", 2)
	q.EnQueue(2)
	fmt.Println("EeQueue: ", 3)
	q.EnQueue(3)


	for _ = range q {
		fmt.Println("DeQueue: ", q.DeQueue())	
	}

}