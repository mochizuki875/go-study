package main

import (
	"io"
	"log"
	"net/http"
	// "time"
)

// チャンネルを構造体のメンバーとして定義
type Queue struct {
	reqCh chan string
}

// メソッドを定義
func (q * Queue) EnQueue(w http.ResponseWriter, r *http.Request) {
	log.Println("[Debug] Before sent")
	q.reqCh <- r.URL.String()
	io.WriteString(w, "Hello world")

	now := <- q.reqCh
	log.Println(now)
}


func main() {
	// Hello world, the web server

	q := Queue{reqCh: make(chan string)}

	// HTTP Handler定義
	// helloHandler := func(w http.ResponseWriter, req *http.Request) {

	// 	log.Println("[Debug] Before sent")

	// 	// reqCh <- timeToString(time.Now())
	// 	reqCh <- 1

	// 	log.Println("[Debug] After sent")

	// 	log.Println("Call helloHandler")
	// 	io.WriteString(w, "Hello, world!\n")
	// 	// io.WriteString(w, "Hello, world!\n")

	// 	now := <- reqCh
	// 	log.Println(now)

	// }

	// Handlerとパスの対応付け
	// http.HandleFunc("/", helloHandler)
	http.HandleFunc("/", q.EnQueue)




	log.Println("[Debug] http.ListenAndServe(\":8080\", nil)")
	// httpサーバー起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}