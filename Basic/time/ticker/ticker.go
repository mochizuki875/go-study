/*
  一定間隔ごとにチャンネルに値を送信し、任意の処理を定期実行させる
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	probeTickerPeriod := time.Duration(5) * time.Second
	probeTicker := time.NewTicker(probeTickerPeriod)

	for {
		select {
		case <-time.NewTicker(time.Duration(rand.Intn(10)+1) * time.Second).C:
			fmt.Println("  [manual]: ", time.Now())
			// probeTicker.Cに値が送信されるタイミングを変更する
			// manualTicker.Cに値が送信された時刻からprobeTicker経過後に値を送信する
			probeTicker.Reset(probeTickerPeriod)
		case <-probeTicker.C:
			fmt.Println("[timer]: ", time.Now())
			// fmt.Printf("%#v\n", t)

		}
	}

}
