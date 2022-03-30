package main
import(
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup // goroutine待機用のWaitGroupを定義（goroutin終了までmainの終了を待機させるため）
	var mu sync. // ★Mutexを定義

	counter := 0

	for i := 0; i < 10000; i++ {

		wg.Add(1) // 実行するgoroutin分wait値を加算
		
		go func(i int) { // 10000回goroutinを実行（並行処理）
			mu.Lock() // ★自分以外の処理をブロック
			defer mu.Unlock() // ★処理が終わったら解放

			counter++ // funcが呼び出される毎にcounterを加算
			fmt.Println("[Debug] go-routin No:", i , "go-routin counter:" , counter) // forループによるgoroutin実行番号を表示

			wg.Done() // goroutin処理が終了したらwait値から-1
			
		}(i)
	}

	wg.Wait() // wait値が0になるまで待機

	fmt.Printf("★main counter = %d\n", counter)
}