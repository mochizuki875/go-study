/*
 ContextのDone()チャンネルをCancelメソッドを使って終了させることで、goroutineを終了させる
*/

package main

import (
	"context"
	"fmt"
	"time"
)

func loop(ctx context.Context) {
	n := 0
	for {
		select {
		case <-ctx.Done(): // Done channelがcloseされたら実行される(ctx.Done()はDone channelを返す)
			return
		default:
			fmt.Println("[goroutine] n = ", n)
			n++
		}
		time.Sleep(time.Second * 1)
	}
}

func main() {

	// Contextとcancel()メソッドを生成
	// ctx.Done() channelを終了したい時はcancel()を実行する
	ctx, cancel := context.WithCancel(context.Background())

	go loop(ctx)
	for i := 0; i < 6; i++ {
		fmt.Println("[main] i = ", i)
		if i == 5 { // i=5ならctx.Done()を終了しbreakする
			fmt.Println("[main] i = 5 close Done Channel by cancel() and break")
			cancel() // 本当はdefer cancel()の方が良いが、ここでは分かりやすいようにこうしている
			break
		}
		time.Sleep(time.Second * 2)
	}

}
