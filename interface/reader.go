package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	r := strings.NewReader("Hello Reader!! I am strings package") // Readerの作成

	b := make([]byte, 8) // len=8(8byte)のbyte型slice

	for { // 無限ループ
		n, err := r.Read(b) // rから8byte分読み取ってbに格納(nには格納されたバイト数が返る)

		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n]) // Readerから値が格納されたbを格納された値のbyte分出力

		fmt.Println("================")

		if err == io.EOF { // EOFまで来たらループを抜ける
			break
		}
	}

}
