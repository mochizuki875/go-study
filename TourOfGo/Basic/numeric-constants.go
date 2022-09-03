package main

import "fmt"

// 数値をconstとして定義する場合自動で方設定をしてくれる
const (
	// 2進数として1を左へ100個ずらす
	Big = 1 << 100

	// 2進数として1を右へ99個ずらす
	Small = Big >> 99
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	// fmt.Println("Big: ", Big, "/Small: ", Small)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	// fmt.Println(needInt(Big))
	fmt.Println(needFloat(Big))
}
