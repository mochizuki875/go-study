package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: // cにxが格納されたら
			fmt.Println("c <- x")
			x, y = y, x+y
		case <-quit: // quit channelから値が取得できたら
			fmt.Println("Quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() { // channelに値が格納されたらその値を取得・出力する処理を別スレッドで起動
		for i := 0; i < 10; i++ {
			fmt.Println("<-c", <-c)
		}
		quit <- 0 // 値を取得し終えたらquitを送信
	}()

	fibonacci(c, quit) // mainスレッドで関数実行
}
