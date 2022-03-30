package main

import "fmt"

// named return value
// 戻り値に名前をつける（x,y）
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return  // naked return
}

func main() {
	fmt.Println(split(17))
}
