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

	// Handlerとパスの対応付け
	http.HandleFunc("/", helloHandler)

	// httpサーバー起動
	log.Println("Server Start")
	http.ListenAndServe(":8080", nil)
}
