package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

var layout = "2006-01-02 15:04:05"
func timeToString(t time.Time) string {
    str := t.Format(layout)
    return str
}

func main() {
	// channel初期化
	q := make(chan string)
	
	//Handler定義
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		log.Println("[Handler] Call helloHandler")
		
		q <- timeToString(time.Now()) // channel送信（timeはリクエスト受信時の値になる）
		io.WriteString(w, "Hello, world!\n")
	}

	// channel受信用goroutine起動
	go func(){
		log.Println("[goroutine] Start goroutine & Receive Chennel")
		for {
			time.Sleep(3*time.Second) // channel受信まで3秒待機
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