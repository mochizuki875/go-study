package main

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

func test_fn() wait.ConditionFunc {
	return func() (done bool, err error) {
		fmt.Println("time: ", time.Now())
		return false, nil
	}
}

func main(){
	const podAttachAndMountRetryInterval time.Duration = 300000000 // 300ms
	const podAttachAndMountTimeout time.Duration = 123000000000 // 2m3s

	fmt.Println("Start time: ", time.Now())
	err := wait.PollImmediate(
		podAttachAndMountRetryInterval,
		podAttachAndMountTimeout,
		test_fn())
	if err != nil {
		fmt.Println("End time: ", time.Now(), " err: ", err)		
	}
}