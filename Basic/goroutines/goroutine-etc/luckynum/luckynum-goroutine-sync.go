package main
import (
	"fmt"
	"time"
	"math/rand"
	"sync"
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
	fmt.Println("[Debug] main() START")

	// sync.WaitGroup構造体を用意
	// 初期カウンタは0
	var wg sync.WaitGroup
	
	// カウンタを+1
	wg.Add(1)

	fmt.Println("what is today's lucky number?")
	go func() {
		defer wg.Done() // goroutine終了時に-1
		getLuckyNum() // 関数呼び出し
	}()

	fmt.Println("[Debug] main() wg.Wait")
	wg.Wait() // wgカウンタが0になるまで終了をブロック

	fmt.Println("[Debug] main() END")
}