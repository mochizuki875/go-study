package main

// 文字列も配列でありsliceを取れるという例
// ここでは文字列末尾の番号を取得する

import "fmt"

func main(){
	Name := "foo-1"
	fmt.Println("Name: ", Name)

	// sliceを取得
	Number := Name[len(Name)-1:]
	fmt.Println("Number: ", Number)
}