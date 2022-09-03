package main

import (
	"fmt"
	"strings"
)

func main() {
	// 2次元の動的配列(slice)を定義
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	fmt.Println(board)
	fmt.Println("board[0] = ", board[0])
	fmt.Println("board[1] = ", board[1])
	fmt.Println("board[2] = ", board[2])

	fmt.Println("=====================")

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
