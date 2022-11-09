package main

import "fmt"

func Summarize(nums []int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

func Diff(x int, y int) int {
	return x - y
}

func main() {
	fmt.Println(Summarize([]int{1, 2, 3, 4}))
}
