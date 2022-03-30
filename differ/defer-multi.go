package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // LIFOの順番で実行
	}
	fmt.Println("done")
}