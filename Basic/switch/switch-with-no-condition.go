package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch { //switchの条件をcaseで指定しているパターン
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
