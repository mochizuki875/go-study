package main

import "fmt"

func echo() string {

	fmt.Println("[echo] Called echo Function")

	// 遅延処理
	// 呼び出し元の関数がreturnする直前に実行
	defer fmt.Println("[echo defer] Echo from Function defer")

	fmt.Println("[echo] Hello!")
	return "[echo] Return Function"
}

func main() {
	fmt.Println(echo())
	fmt.Println("Echo From main")

}
