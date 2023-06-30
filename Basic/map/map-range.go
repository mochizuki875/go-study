/*
	mapをrangeを使ってforで回す例
	必ずしも先頭から評価されるわけじゃ無いので注意
*/

package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["A"] = 1
	m["B"] = 2
	m["C"] = 3

	for i := 0; i < 100; i++ {
		fmt.Printf("\n === Start Loop ===\n")
		for key, value := range m {
			fmt.Println(key, value)
		}
	}

}
