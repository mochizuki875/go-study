package main

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

func main(){
	const backOffPeriod time.Duration = 10000000000 // 10s
	const workerBackOffPeriodJitterFactor = 0.5

	for {
		t := wait.Jitter(backOffPeriod, workerBackOffPeriodJitterFactor)
		fmt.Println(t)
		time.Sleep(t)
	}
}