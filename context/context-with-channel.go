package main

import (
	"context"
	"fmt"
)

// 受信channelを戻り値とする(<-chan int)
func gen(ctx context.Context) <-chan int {
	dst := make(chan int)

	go num(ctx, dst) // 別スレッドでnum()を実行

	fmt.Println("return dst as channel")
	return dst // dst channelを受信channelとして戻り値として返す
}

func num(ctx context.Context, dst chan int) {
	n := 1
	for { // 無限ループ
		select {
		case <-ctx.Done(): // Done channelがcloseされたら実行される(ctx.Done()はDone channelを返す)
			return // returning not to leak the goroutine
		case dst <- n: // dstに送信可能な時はdst <- nが実行される
			fmt.Println("dst <- ", n)
			n++
		}
	}
}

func main() {

	ctx, cancel := context.WithCancel(context.Background()) // ★Done channelを持つcontextとcancel関数を返す

	defer fmt.Println("main() called cancel(), receive <-ctx.Done()")
	defer cancel() // cancelが呼ばれるとDone channelがcloseされる

	// genを実行
	// gen(ctx)の戻り値は<-chan intなのでchannelが受信した値がnに入る
	for n := range gen(ctx) { // channelがcloseするまで受信を続ける
		fmt.Println(n)
		if n == 5 { // gen channelに5が格納されるとbreakしてcancel()が呼ばれる
			fmt.Println("n=5 break")
			break
		}
	}
}
