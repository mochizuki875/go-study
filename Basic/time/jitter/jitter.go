package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/wait"
	"time"
)

func main() {

	for {
		// Jitter returns a time.Duration between duration and duration + maxFactor * duration.
		// 10s~15sの間をランダム
		t := wait.Jitter(time.Second*10, 0.5)
		fmt.Println(t)
		time.Sleep(t)
	}
}
