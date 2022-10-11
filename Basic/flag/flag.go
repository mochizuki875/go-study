package main

import (
	"flag"
	"fmt"
)

func main() {

	/*
		フラグはコマンドライン引数を扱うパッケージ
		https://pkg.go.dev/flag

		ここでは-sと言うコマンドライン引数を扱えるようにする
		  go run flag.go -s bbb
		  →bbb
	*/

	// フラグの値を格納するポインタとして返される
	// フラグの名前、デフォルト値、使い方の説明(-sで表示される)
	config := flag.String("s", "aaa", "文字列flagの例")

	// コマンドライン引数で渡された値をパースする
	flag.Parse()

	fmt.Printf("type: %t, Value: %v\n", config, config)
	fmt.Println(*config) // configはポインタなので*configとしてポインタの示す値を取得

}
