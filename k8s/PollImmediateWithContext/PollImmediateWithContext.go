package main

import (
	"context"
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

func test_fn() wait.ConditionWithContextFunc {
	return func(ctx context.Context) (done bool, err error) {
		fmt.Println("time: ", time.Now())
		return false, nil
	}
}

func main(){
	const podAttachAndMountRetryInterval time.Duration = 300000000 // 300ms
	const podAttachAndMountTimeout time.Duration = 63000000000 // 1m3s
	
	ctx, cancel := context.WithCancel(context.Background())

	// 10秒後にcancel()を実行しctx.Done() channelを終了させる
	go func(cancel context.CancelFunc) {
		time.Sleep(time.Second * 10)
		cancel()
	}(cancel)

	fmt.Println("Start time: ", time.Now())

	// 300ms間隔で1m3s間test_fn()を実行するループ(→これをcancel()で止める)
	err := wait.PollImmediateWithContext(
		ctx,
		podAttachAndMountRetryInterval,
		podAttachAndMountTimeout,
		test_fn())
	if err != nil {
		fmt.Println("End time: ", time.Now(), " err: ", err) // タイムアウトの場合はエラーになる	
	}
}