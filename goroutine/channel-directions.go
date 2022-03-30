package main
import(
	"fmt"
)

// msgをchannelに送信
func ping(pings chan<- string, msg string){
	pings <- msg
}

// pings channelの値をpongs channelに渡す
func pong(pings <-chan string, pongs chan<- string){
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "Hello")
	pong(pings, pongs)

	fmt.Println(<-pongs) // pongs channelから受信した値を表示

}