package main
import (
	"fmt"
	"time"
	"math/rand"
)

func getLuckyNum(c chan <- int){
	fmt.Println("...")

	// 一定時間待機
	rand.Seed(time.Now().Unix())
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

	num := rand.Intn(10)
	c <- num // channelにnumを送信

}

func main() {
	fmt.Println("[Debug] main() START")

	c := make(chan int) // channelの作成

	fmt.Println("what is today's lucky number?")

	go getLuckyNum(c) // goroutine

	// channelから値を受信
	// 値を受信するまで待機
	num := <- c 

	// goroutineからラッキーナンバーを取得して表示
	fmt.Printf("Today's your lucky number is %d\n", num)

	close(c) //channelをclose

	fmt.Println("[Debug] main() END")
}