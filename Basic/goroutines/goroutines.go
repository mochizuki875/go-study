package main

// Go Routineはユーザースレッドによる並行処理

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i, "回目: ", s)
	}
}

func main() {
	go say("Hello From goroutine") // ユーザースレッドを新規で作成して処理を実行(go routine)
	say("Hello from main routine")
}
