/*
 ContextをWithDeadlineで生成することで、Done()チャンネルを時限式で終了させる。
 これに伴いgoroutineを終了させる。
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
		case <-ctx.Done(): // Done channelがcloseされたら実行される(Contextが生成されてから3秒後にclose)
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
	// 生成してから3秒後にDeadlineを迎えDone()チャンネルがクローズされる
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))

	defer cancel()

	go loop(ctx)
	for i := 0; i < 6; i++ {
		fmt.Println("[main] i = ", i)
		if i == 5 { // i=5ならctx.Done()を終了しbreakする
			fmt.Println("[main] i = 5 break")
			break
		}
		time.Sleep(time.Second * 2)
	}

}
