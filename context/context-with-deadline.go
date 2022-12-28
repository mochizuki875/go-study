package main

import (
	"context"
	"fmt"
	"time"
)

// const shortDuration = 10 * time.Second // 10s後にDone channelをclose

const shortDuration = 1 * time.Millisecond // 1ms後にDone channelをclose

func main() {

	// dをDeadlineとするcontextを返す
	// Deadlineを過ぎる、またはcancel()が呼ばれるとDone channelがcloseされる
	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel() // Done channelをclose

	select {
	case t := <-time.After(1 * time.Second): // channelがtrueの時(1s経過後にchannelが時刻を受信する)
		fmt.Println("over slept. (t = ", t, ")")
		fmt.Println(ctx.Err())
	case <-ctx.Done(): // channelがtrueの時
		fmt.Println("ctx.Done() is true. Done channel is already closed.")
		fmt.Println(ctx.Err())
	}
}
