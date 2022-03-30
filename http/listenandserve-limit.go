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

	// HTTP Handler定義
	helloHandler := func(w http.ResponseWriter, req *http.Request) {


		log.Println("Call helloHandler")
		io.WriteString(w, "Hello, world!\n")

	}
	// Handlerとパスの対応付け
	http.HandleFunc("/", helloHandler)

	max-connections := 10
	limit_listener := netutil.LimitListener(listener, c.Int("max-connections"))

	defer limit_listener.Close()

	log.Println("[Debug] http.ListenAndServe(\":8080\", nil)")



	// httpサーバー起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}