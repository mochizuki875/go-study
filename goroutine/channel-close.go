package main
import (
	"fmt"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool) // goroutineの完了をmainに伝えるためのchannel

	go func(){
		for {
			j, more := <-jobs
			fmt.Println("[Debug] more =", more)
			if more {
				fmt.Println("received job", j)
			} else { 
				fmt.Println("received all jobs")
				done <- true // 全てのjobsを受信したらdonechannelにtrueを送信
				return
			}
		}
	}()

	for j := 1; j <= 5; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	// jobs channelをclose
	// これによりmore=falseになる
	close(jobs)
	fmt.Println("sent all jobs")

	// doneチャンネルの受信待ち(mainはchannelの受信が完了するまで待機する)
	// mainをgo channelが閉じられるまで待機する目的なので受信先は指定しない
	<- done
	close(done)
	
}