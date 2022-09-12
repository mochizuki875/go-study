package main
import (
	"fmt"
	"time"
	"sync"
)

func main() {
	fmt.Println("[Debug] main() START")

	// sync.WaitGroup構造体を用意
	// 初期カウンタは0
	var wg sync.WaitGroup
	
	// カウンタを+1
	wg.Add(1)
	fmt.Println("[Debug] wg.Add(1)")

	go func() {
		defer wg.Done() // goroutine終了時に-1
		defer fmt.Println("[Debug] goroutine wg.Done()")
		fmt.Println("[Debug] goroutine 5 sec sleep")
		time.Sleep(5*time.Second)
	}()

	fmt.Println("[Debug] main() wg.Wait")
	// wgカウンタが0になるまで終了をブロック
	// 待ち合わせを入れないとgoroutinが終わる前にmainが終了してしまう
	wg.Wait()

	fmt.Println("[Debug] main() END")
}