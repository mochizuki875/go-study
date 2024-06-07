package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	hellopb "grpc-learning/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GreetingServiceServerインターフェースを実装したstruct
type myServer struct {
	hellopb.UnimplementedGreetingServiceServer
}

func NewMyServer() *myServer {
	return &myServer{}
}

func (s *myServer) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	// リクエストからnameフィールドを取り出して
	// "Hello, [名前]!"というレスポンスを返す
	// return &hellopb.HelloResponse{}, fmt.Errorf("response error") // このエラーもclientにレスポンスとして返る
	return &hellopb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}

func main() {
	// 1. 8080番portのLisnterを作成
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// 2. gRPCサーバーを作成
	// s := grpc.NewServer()
	// rf. https://github.com/kubernetes/kubernetes/blob/e67f889edc4ab278028f6cffd2501bc90a0defcf/pkg/kubelet/apis/podresources/client.go#L48
	s := grpc.NewServer(grpc.MaxSendMsgSize(15), // 送信するメッセージの最大サイズを設定
		grpc.MaxRecvMsgSize(20)) // 受信するメッセージの最大サイズを設定

	// 3. gRPCサーバーにGreetingServiceを登録
	hellopb.RegisterGreetingServiceServer(s, NewMyServer())

	// 4. 作成したgRPCサーバーを、8080番ポートで稼働させる
	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()

	// [デバッグ用] gRPCurl用にサーバーリフレクションの設定
	reflection.Register(s)

	// 5.Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
