package main
import (
	"fmt"
	"time"
	"math/rand"
)

func getLuckyNum(){
	fmt.Println("...")

	// 一定時間待機
	rand.Seed(time.Now().Unix())
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

	num := rand.Intn(10)
	fmt.Printf("Today's your lucky number is %d\n", num)
}

func main() {

	fmt.Println("main() START")

	fmt.Println("what is today's lucky number?")
	go getLuckyNum()

	time.Sleep(time.Second * 5) // goroutineの終了を待つ

	fmt.Println("main() END")
}