package main
import(
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup // goroutine待機用のWaitGroupを定義（goroutin終了までmainの終了を待機させるため）

	counter := 0

	for i := 0; i < 10000; i++ {

		wg.Add(1) // 実行するgoroutin分wait値を加算
		
		go func(i int) { // 10000回goroutinを実行（並行処理）

			counter++ // funcが呼び出される毎にcounterを加算 → goroutinは並行処理なので、加算処理も並行で行われてしまうことがある
			fmt.Println("[Debug] go-routin No:", i , "go-routin counter:" , counter) // forループによるgoroutin実行番号を表示

			wg.Done() // goroutin処理が終了したらwait値から-1
			
		}(i)
	}

	wg.Wait() // wait値が0になるまで待機

	fmt.Printf("★main counter = %d\n", counter)
}