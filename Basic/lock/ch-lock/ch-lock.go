/*
channelを使って処理をブロックする
*/

package main

import "fmt"

func main() {
	stopCh := make(chan int)

	go func(stopCh chan int) {
		defer func() { stopCh <- 0 }() // channelに値を送信してmain routineのブロック解除

		for i := 0; i < 100; i++ {
			fmt.Println("Count = ", i)
		}
	}(stopCh)

	// channelを使ってmain routineの終了をブロック
	fmt.Println("stopCh = ", <-stopCh)

}
