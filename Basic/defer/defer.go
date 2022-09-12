package main

import "fmt"

func main() {
	// 遅延処理
	// 呼び出し元の関数がreturnする時に実行
	defer fmt.Println("world")

	fmt.Println("Hello")
}
