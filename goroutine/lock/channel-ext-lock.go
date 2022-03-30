package main
import (
	"fmt"
	"sync"
)

type DoubleCounter struct {
	valueA int
	valueB int
	commandCh chan int
	wg sync.WaitGroup
}

func NewDoubleCounter() *DoubleCounter {
	dc := &DoubleCounter{
		valueA: 0,
		valueB: 0,
		commandCh: make(chan int),
	}

	go dc.Start()
	return dc
}

// channel受信用の無限ループ
// 受信したら加算処理を実施
func (dc *DoubleCounter) Start() {
	for {
		target := <- dc.commandCh
		fmt.Println("[Debug] Receive channel in loop:", target)
		if target %2 == 0 {
			dc.valueA++
		} else {
			dc.valueB++
		}
		dc.wg.Done()
	}
}

func (dc *DoubleCounter) Sent(t int) {
	fmt.Println("[Debug] Send channel:", t)
	dc.commandCh <- t
}

func (dc *DoubleCounter) Wait() {
	dc.wg.Wait()
}

func (dc *DoubleCounter) Print() {
	fmt.Printf("A: %d, B: %d\n", dc.valueA, dc.valueB)
}

func main(){
	dc := NewDoubleCounter()
	for i:=0; i < 1000; i++ {

		dc.wg.Add(1)
		go func(t int) {
			dc.Sent(t)
		}(i) // goroutinで実行する無名関数に渡す引数（tに該当）
	}

	dc.Wait()
	dc.Print()

}