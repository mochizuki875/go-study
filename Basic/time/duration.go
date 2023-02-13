/*
  sleepのdurationを倍々で増加させていく
  最大で5s
*/

package main

import (
	"fmt"
	"math"
	"time"
)

func main(){
	const (
		base   = 100 * time.Millisecond
		max    = 5 * time.Second
		factor = 2
	)
	duration := base

	for {
		// exponential backoff
		time.Sleep(duration)
		duration = time.Duration(math.Min(float64(max), factor*float64(duration)))
		fmt.Println("duration: ", duration)
	}
}
