package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i < 10; i++ { // 条件句に()不要
		sum += i
		fmt.Printf("i = %v , sum = %v\n", i, sum)
	}
}
