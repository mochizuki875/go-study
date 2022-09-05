package main

import (
	"context"
	"fmt"
)

func main() {

	// channelを戻り値とし、関数内のgoroutineでそこに値を送信する
	gen := func(ctx context.Context) <-chan int { // int型のchannelを戻り値とするgenを定義(gen自体がchannelの受信になる)
		dst := make(chan int)
		n := 1
		go func() { // 別スレッドで実行
			for {
				select {
				case <-ctx.Done(): // Done channelがtrueの時
					return
				case dst <- n: // dstに送信可能な時はdst <- nが実行される
					fmt.Println("dst <- n")
					n++
				}
			}
		}()
		fmt.Println("return dst as channel")
		return dst // dst channelに送信された値が戻り値
	}

	ctx, cancel := context.WithCancel(context.Background()) // Done channelを持つcontextとcancel関数を返す

	defer cancel() // cancelが呼ばれるとDone channelがクローズされる(ctx.Done())

	// genを実行
	// gen(ctx)の戻り値は<-chan intなのでchannelが受信した値が順次nに入る
	for n := range gen(ctx) { // channelがcloseするまで受信を続ける
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
