package main
import(
	"fmt"
)

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3 // DeadLock

	// channel bufferについてもlenとcapを取得可能
	fmt.Println("len(ch) = ", len(ch), "cap(ch)= ", cap(ch))
	chLen := len(ch)

	for i := 0; i < chLen; i++ {
		fmt.Println(<-ch) // channelの受信（FIFO）
		fmt.Println("len(ch) = ", len(ch), "cap(ch)= ", cap(ch))
	}

}