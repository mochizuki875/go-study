package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // channelに値を送信し終えたのでclose(受信側がcloseを知る必要がある時だけ明示的にcloseする)
}

func main() {
	c := make(chan int, 10)

	go fibonacci(cap(c), c)

	for i := range c { // channelがcloseするまで受信を続ける
		fmt.Println(i)
	}

}
