package main

import (
	"fmt"
)

var i, j int = 1, 2

func main() {
	// 型名を省略可能
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
