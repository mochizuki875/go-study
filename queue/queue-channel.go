/*
Goでキューを扱いたい
スレッド間でキューを使ってやりとりできるようにしたい
各スレッドで扱うシーケンスA,Bを順序性を持たせて実行する
キューをtypeで定義し、キューに対する処理をメソッドとして実装
*/

package main
import (
	"fmt"
	"time"
)

type queue chan int

func CreateQueue() queue {
	q := make(chan int) 
	return q
}

// キューに値を追加するメソッド
func (q *queue) EnQueue(t int) {
	fmt.Println("[Enqueue Routine] Sequence-A:", t)
	(*q) <- t
}

// キューから値を取り出すメソッド
func (q *queue) DeQueue() {
	t := <- *q
	fmt.Println("[Dequeue Routine] Sequence-B:", t)
}

func main() {
	
	q := CreateQueue()

	// 3秒間隔でキューに値を格納するスレッド
	go func(){
		for i:=0; i < 10; i++ {
			q.EnQueue(i) 
		}
	}()

	// キューに値が格納されたら取り出して表示するスレッド
	// 無限ループだけどメインルーチン終了時に終了される
	go func(){
		for {
			q.DeQueue()
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("[main] sleep:", i+1)
		time.Sleep(1*time.Second)
	}
}