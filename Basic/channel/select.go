package main

import (
	"fmt"
	"time"
)

func main() {

	msgChA := make(chan string)
	msgChB := make(chan string)
	var receiveA string
	var receiveB string

	probeTickerA := time.NewTicker(time.Duration(3) * time.Second)
	probeTickerB := time.NewTicker(time.Duration(10) * time.Second)

	defer probeTickerA.Stop()
	defer probeTickerB.Stop()

	// 一定間隔でチャンネルに値を送信するgoroutine
	go func() {
		for {
			select { // selectを使うことで各チャンネルに値が送信されるまでは処理をブロック
			case <-probeTickerA.C:
				msgChA <- "message from A"
			case <-probeTickerB.C:
				msgChA <- "message from B"
			}
		}
	}()

	// チャンネルから値を受け取って表示
	for {
		select {
		case receiveA = <-msgChA:
			fmt.Println("received message: ", receiveA)
		case receiveB = <-msgChB:
			fmt.Println("received message: ", receiveB)
		default: // msgChA, msgChBのどちらにも値がない場合、処理をブロックせずにdefaultとして定義した処理を行う
			// non-blocking-channel-operation
			fmt.Println("no message received(non-blocking-channel-operation)")
		}
		// 1秒間スリープ
		time.Sleep(1 * time.Second)
	}

}
