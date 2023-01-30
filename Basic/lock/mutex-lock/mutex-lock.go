/*
mutexを使って処理をブロックする
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex

	go func() { // goroutineでユーザースレッド分離
		defer fmt.Println("goroutine_1がロックを解除")
		defer mu.Unlock() // ロック解除(排他制御を解除)
		mu.Lock()         // ロック(排他制御により自分以外のユーザースレッド(main,routine2)の処理をロックする) 俺だけに処理させて他は止まれ!のイメージ
		fmt.Println("goroutine_1がロック")
		for i := 0; i < 5; i++ {
			fmt.Println("[goroutine_1]Count = ", i)
		}
	}()

	go func() { // goroutineでユーザースレッド分離
		defer fmt.Println("goroutine_2がロックを解除")
		defer mu.Unlock() // ロック解除(排他制御を解除しmain routinの処理を継続)
		mu.Lock()         // ロック(排他制御により自分以外のユーザースレッド(main,routine2)の処理をロックする) 俺だけに処理させて他は止まれ!のイメージ
		fmt.Println("goroutine_2がロック")
		for i := 0; i < 10; i++ {
			fmt.Println("[goroutine_2]Count = ", i)
		}
	}()

	fmt.Println("goroutineがmu.Lockするための時間を待機")
	time.Sleep(5 * time.Second)
}
