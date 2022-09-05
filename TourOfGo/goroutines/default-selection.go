package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	fmt.Printf("tick type is %t\n", tick)
	boom := time.After(500 * time.Millisecond)
	fmt.Printf("boom type is %t\n", boom)

	fmt.Println("========================")

	for { // 無限ループ
		select {
		case t := <-tick: // tick channelに値が入っている場合(v, ok := <- tickでok=trueの場合)
			fmt.Println("tick", t)
		case b := <-boom: // boom channelに値が入っている場合(v, ok := <- boomでok=trueの場合)
			fmt.Println("Boom", b)
			return
		default:
			fmt.Println("     .")
			time.Sleep(50 * time.Millisecond)

		}
	}
}
