/*

Goのhttp.Handlerやhttp.HandlerFuncをちゃんと理解する
https://journal.lampetty.net/entry/understanding-http-handler-in-go
*/

package main

import (
	"io"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	log.Println("Call helloHandler")
	io.WriteString(w, "Hello, world!\n")
}

func main() {

	// ServeMuxを作成
	mux := http.NewServeMux()

	/*
		Handleメソッドでパスと対応するHandler(ServeHTTPを含むInterface)を登録する
		第2引数はHandlerに対してHandlerFuncという関数型に任意の関数をキャストして代入(実装)
		→HandlerFunc型にはServeHTTPメソッドが実装されている

		最終的にmux.mにパスとHandlerの組み合わせが登録される
	*/
	mux.Handle("/hello", http.HandlerFunc(hello))

	// 指定したポートでリッスンするhttpサーバーを起動
	// muxに登録されたHandlerを実行
	http.ListenAndServe(":8080", mux)

}
