package main

import "fmt"

func main() {
	switch x := 2; x { // xの値によってcaseを定義
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println(x)
	}
}
