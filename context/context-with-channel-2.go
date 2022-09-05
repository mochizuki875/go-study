package main

import (
	"context"
	"fmt"
)

// channelが受信した値を戻り値とする(<-chan int)
func gen(ctx context.Context) <-chan int {
	dst := make(chan int)

	go num(ctx, dst) // 別スレッドでnum()を実行

	fmt.Println("return dst as channel")
	return dst // dst channelの受信を戻り値として返す
}

func num(ctx context.Context, dst chan int) {
	n := 1
	for { // 無限ループ
		select {
		case <-ctx.Done(): // Done channelがtrueの時(main channelでcancel()が実行されてDone channelがcloseし、ctx.Done()がtrueとなる)
			return // returning not to leak the goroutine
		case dst <- n: // dstに送信可能な時はdst <- nが実行される
			fmt.Println("dst <- n")
			n++
		}
	}
}

func main() {

	ctx, cancel := context.WithCancel(context.Background()) // Done channelを持つcontextとcancel関数を返す

	defer cancel() // cancelが呼ばれるとDone channelがクローズされる(→ctx.Done()に値が送信されtrueになる)

	// genを実行
	// gen(ctx)の戻り値は<-chan intなのでchannelが受信した値がnに入る
	for n := range gen(ctx) { // channelがcloseするまで受信を続ける
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
