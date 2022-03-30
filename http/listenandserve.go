package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server

	// HTTP Handler定義
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		log.Println("Call helloHandler")
		io.WriteString(w, "Hello, world!\n")
	}

	log.Println("[Debug] http.ListenAndServe(\":8080\", nil)")

	Handlerとパスの対応付け
	http.HandleFunc("/", helloHandler)

	// httpサーバー起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}