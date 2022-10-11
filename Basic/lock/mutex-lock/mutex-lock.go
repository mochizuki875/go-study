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
		defer mu.Unlock() // ロック解除(排他制御を解除しmain routinの処理を継続)
		mu.Lock()         // ロック(排他制御によりmain routinの処理をロックする) 俺だけに処理させて他は止まれ!のイメージ
		for i := 0; i < 100; i++ {
			fmt.Println("Count = ", i)
		}
	}()

	fmt.Println("goroutineがmu.Lockするための時間を待機")
	time.Sleep(3 * time.Second)
}
