/*
  continue: ループ内で今のループ処理を終えて次のループ処理に入る
  3の倍数でcontinueを利用
*/

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%3 == 0 && i != 0 {
			fmt.Println("ヒャー!")
			continue
		}
		fmt.Println("i = ", i)
	}
}