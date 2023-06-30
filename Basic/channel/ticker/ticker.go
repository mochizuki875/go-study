/*
  一定間隔ごとにチャンネルに値を送信し、任意の処理を定期実行させる
*/

package main

import (
	"fmt"
	"time"
)

func main() {

	probeTickerPeriod := time.Duration(30) * time.Second
	probeTicker := time.NewTicker(probeTickerPeriod)

	for {
		select {
		case t := <-probeTicker.C:
			// fmt.Println(time.Now())
			fmt.Println(t)

		}
	}

}
