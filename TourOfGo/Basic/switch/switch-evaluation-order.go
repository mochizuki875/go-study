package main

import (
	"fmt"
	"time"
)

func main() {

	// 日曜日が今日から何日後かを表示

	today := time.Now().Weekday() // 今日の曜日を取得
	fmt.Println("Today is", today)

	fmt.Println("When's Saturday?")
	switch time.Saturday { // time.Saturday = case となるcaseを上から順に評価
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
