/*
WaitGroupを使って処理をブロックする
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1) // ロック(WGにAddされた回数Doneされるまでロック)

	go func() {
		defer wg.Done() // ロック解除
		for i := 0; i < 100; i++ {
			fmt.Println("Count = ", i)
		}
	}()

	wg.Wait() // 待機

}
