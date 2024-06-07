/*
condを使って処理をブロックする
Signalを使って処理を開始させる
https://selfnote.work/20220718/programming/golang-concurrent-signal/
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var queue []string
	lock := sync.NewCond(&sync.Mutex{})

	go func() { // goroutineでユーザースレッド分離
		defer fmt.Println("[goroutine] ロックを解除") // 読み込みのための排他制御を解除
		defer lock.L.Unlock()

		fmt.Println("[goroutine] ロック") // 読み込みのための排他制御
		lock.L.Lock()

		// シグナルの受信を待機
		for len(queue) == 0 {
			fmt.Println("[goroutine] シグナルの受信を待機")
			lock.Wait()
		}

		// シグナルを受信してUnlockされると以下の処理が実行される

		fmt.Println("[goroutine] シグナルを受信")
		fmt.Println("[goroutine] queue = ", queue[0])
		queue = queue[1:]

	}()

	time.Sleep(3 * time.Second) // goroutineのシグナル待機が開始するまで待機

	fmt.Println("ロック") // 書き込みのための排他制御
	lock.L.Lock()

	// queueにメッセージを追加してシグナルを送信
	fmt.Println("queueにメッセージを追加")
	queue = append(queue, "message")
	fmt.Println("シグナルを送信")
	lock.Signal()

	fmt.Println("ロックを解除") // 書き込みのための排他制御を解除
	lock.L.Unlock()

	time.Sleep(5 * time.Second)

}
