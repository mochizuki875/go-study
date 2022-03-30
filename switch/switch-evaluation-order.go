package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now().Weekday()
	fmt.Println("Today is", today)
	fmt.Println("When's Saturday?")
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In Two Days.")
	default:
		fmt.Println("Too far away.")
	}
}
