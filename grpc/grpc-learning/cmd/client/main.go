package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	hellopb "grpc-learning/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	scanner *bufio.Scanner
	client  hellopb.GreetingServiceClient
)

func main() {
	fmt.Println("start gRPC Client.")

	// 1. 標準入力から文字列を受け取るスキャナを用意
	scanner = bufio.NewScanner(os.Stdin)

	// 2. gRPCサーバーとのコネクションを確立
	address := "localhost:8080"

	// 非推奨
	// conn, err := grpc.Dial(
	// 	address,
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// 	grpc.WithBlock(),
	// )

	conn, err := grpc.DialContext(context.TODO(), address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(20)), // 送信するメッセージの最大サイズを設定
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(30))) // 受信するメッセージの最大サイズを設定

	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	// 3. gRPCクライアントを生成
	client = hellopb.NewGreetingServiceClient(conn)

	for {
		fmt.Println("1: send Request")
		fmt.Println("2: exit")
		fmt.Print("please enter >")

		scanner.Scan()
		in := scanner.Text()

		switch in {
		case "1":
			grpcHello()

		case "2":
			fmt.Println("bye.")
			goto M
		}
	}
M:
}

func grpcHello() {
	fmt.Println("Please enter your name.")
	scanner.Scan()
	name := scanner.Text()

	req := &hellopb.HelloRequest{
		Name: name,
	}
	res, err := client.Hello(context.Background(), req) // GreetingServiceClientに実装されているHelloメソッドを呼び出す
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.GetMessage())
	}
}
