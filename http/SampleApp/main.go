package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// ハンドラー関数を定義する
	handler1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello Go\n")
	}

	// パスとハンドラー関数を結びつける
	http.HandleFunc("/", handler1)

	// ポート番号 8080 で待ち受けを開始
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
