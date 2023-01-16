package main

import (
	"fmt"
	"os"
)

func main() {
	uid := os.Getuid()
	fmt.Println(uid)
	if uid == 0 {
		fmt.Println("rootユーザーです。")
	}
}
