package main
import(
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 1)
	
	go func() {
		time.Sleep(3*time.Second)
		c1 <- "result 1"
	}()

	select{ // 最初に受信したものを処理する
	// このケースではtime.Afterが先に受信する
	case res := <- c1: // c1から値を受信した場合
		fmt.Println(res)
	case <-time.After(1 * time.Second): // time.Afterを受信した場合
		fmt.Println("timeout 1")
	}
}