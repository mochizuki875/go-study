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
	// Hello world, the web server

	// queueになるチャンネルを定義
	reqCh := make(chan string) 

	// HTTP Handler定義
	helloHandler := func(w http.ResponseWriter, req *http.Request) {

		log.Println("[Debug] Before sent")

		reqCh <- timeToString(time.Now()) // ここで止まる？

		log.Println("[Debug] After sent")

		log.Println("Call helloHandler")
		io.WriteString(w, "Hello, world!\n")
		// io.WriteString(w, "Hello, world!\n")

		now := <- reqCh
		log.Println(now)

	}

	// Handlerとパスの対応付け
	http.HandleFunc("/", helloHandler)

	log.Println("[Debug] http.ListenAndServe(\":8080\", nil)")
	// httpサーバー起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}