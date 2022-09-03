package main

import "fmt"

func main() {
	sum := 1
	for sum < 100 { // 100未満にする
		if (sum + sum) >= 100 {
			break
		} else {
			sum += sum
		}
		fmt.Println(sum)
		if sum > 100 {
			break
		}
	}
	fmt.Println("===============")
	fmt.Println("sum = ", sum)
}
