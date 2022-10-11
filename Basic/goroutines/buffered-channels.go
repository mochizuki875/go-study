package main

import "fmt"

func main() {
	const N = 20
	ch := make(chan int, N) // 長さNのchannel

	// channelに値を送信
	for i := 0; i < N; i++ {
		ch <- i
	}

	// channelから値を受信(FIFO)
	for i := 0; i < N; i++ {
		fmt.Println(<-ch)
	}
}
