/*
チャンネルにrangeを使うと受信待機してチャンネルが受信した値を随時取得できる。
*/

package main

import (
	"fmt"
	"time"
)

// チャネルに値を渡す①
func sendValue_1(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
}

// チャネルに値を渡す②
func sendValue_2(c chan<- int) {
	for i := 5; i < 10; i++ {
		c <- i
	}
}

// チャネルから値を受け取る
func receiveValue(c <-chan int) {
	for v := range c { // チャンネルの受信を待機
		fmt.Println("チャネルから受け取った値：", v)
	}
}

func main() {
	c := make(chan int)
	
	go sendValue_1(c) // チャンネルに値を送信①
	go receiveValue(c) // チャンネルの値を受け取って表示するgoroutine(closeされるまで待機する)
	time.Sleep(3 * time.Second)
	go sendValue_2(c) // チャンネルに値を送信②

	time.Sleep(10 * time.Second)
	close(c)

	fmt.Println("処理を終了します。")
}
