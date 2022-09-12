package main

import "fmt"

func main() {
	sum := 1
	for sum < 100 { // 条件を満たす(true)である限り処理を継続(100を超えたらループを抜ける)
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println("===============")
	fmt.Println("sum = ", sum)
}
