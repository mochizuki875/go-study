/*
httpリクエスト時に呼び出されるhandlerの中でchannel送信を実施することにより、
リクエスト時に生成されたスレッドのバッファリング（直列処理）を行う
*/

package main

import (
	"io"
	"log"
	"net/http"
	"time"
	"sync"
)

// timeを所定フォーマット文字列に変換する関数
var layout = "2006-01-02 15:04:05"
func timeToString(t time.Time) string {
    str := t.Format(layout)
    return str
}

func main() {
	// channel初期化
	q := make(chan string)
	var wg sync.WaitGroup // ★wait

	//Handler定義
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		log.Println("[Handler] Call helloHandler")

		wg.Add(1) // ★wait コレだとスレッド数分一気に入っちゃう
		defer wg.Done() wg.Wait() // ★wait
		defer fmt.Println("[Debug] goroutine wg.Done()") // ★wait
		
		q <- timeToString(time.Now()) // channel送信（timeはリクエスト受信時の値になる）
		io.WriteString(w, "Hello, world!\n")
	}

	// channel受信用goroutine起動
	go func(){
		log.Println("[goroutine] Start goroutine & Receive Chennel")
		for {
			// channel受信まで3秒待機
			// 本当はスレッド処理終了を検知したい（waitとか使えないかな？）
			wg.Wait() // ★wait
			time.Sleep(3*time.Second) 

			msg := <-q
			log.Println("[goroutine]Receive Channel:", msg)
		}
	}()

	log.Println("[main] http.ListenAndServe(\":8080\", nil)")

	// Handlerとパスの対応付け
	http.HandleFunc("/", helloHandler)

	// httpサーバー起動
	log.Fatal(http.ListenAndServe(":8080", nil))


}