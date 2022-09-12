package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	c <- sum // channelにsumを入れる
}

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{6, 7, 8, 9}

	c := make(chan int) // int型のchanelを定義

	go sum(s1, c)
	go sum(s2, c)
	go sum(s1[:len(s1)/2], c)

	// channelから値を受信して変数に格納
	// channelの受信が終わるまでここで待機(=goroutineが完了するまでmainスレッドを待機)
	x, y, z := <-c, <-c, <-c
	fmt.Println(x, y, z)

}
