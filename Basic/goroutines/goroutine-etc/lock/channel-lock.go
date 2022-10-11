package main
import(
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup // goroutine待機用のWaitGroupを定義（goroutin終了までmainの終了を待機させるため）
	ch := make(chan int) // channel作成

	counter := 0

	for i := 0; i < 1000; i++ {

		wg.Add(1) // 実行するgoroutin分wait値を加算
		
		go func(t int) { // 10000回goroutinを実行（並行処理）
			
			fmt.Println("[Debug] Send channel:", t)
			ch <- t // channelに値を送信
			
		}(i) // goroutinで実行する無名関数に渡す引数（tに該当）
	}

	// 無限ループ関数をgoroutinで回す
	// channelから値を受信したらcounterを加算
	// waitカウンタが0になりmainルーチンが終了すると同時に終了
	go func() {
		for {
			i := <- ch // channelから値を受信するまで待機

			counter++
			fmt.Println("[Debug] go-routin No:", i , "go-routin counter:" , counter) // forループによるgoroutin実行番号を表示	
			wg.Done() // 加算処理1回ごとにwaitカウンタから-1
		}
	}()

	wg.Wait() // wait値が0になるまで待機

	fmt.Printf("★main counter = %d\n", counter)
}