# gRPC

作ってわかる！ はじめてのgRPC  
https://zenn.dev/hsaki/books/golang-grpc-starting

gRPCに入門しよう！  
https://qiita.com/Ian_C/items/523912ad9b0169cacb3c

protocのインストール  
https://protobuf.dev/


## 大まかな流れ
- `.proto`ファイルを作成する
  - サービスやメソッド、型を定義する
- `.proto`ファイルから`protoc`コマンドを使ってGoのコードを生成する
  - サービスのインターフェイス定義や型定義（struct）、コンストラクタなどが生成される
- サーバーを作成する
  - サービスのインターフェイスを実装してgRPCサーバーに登録する（登録用のメソッドも自動生成される）
  - gRPCサーバーを起動する
- クライアントを作成する
  - gRPCサーバーとのコネクションとクライアントのコンストラクタを使ってgRPCクライアントを作成する


```bash
curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v27.0/protoc-27.0-linux-x86_64.zip
mkdir protoc-27.0-linux-x86_64
unzip protoc-27.0-linux-x86_64.zip -d protoc-27.0-linux-x86_64
sudo mv ./protoc-27.0-linux-x86_64/bin/protoc /usr/bin/protoc 
```

Protocol Buffers Compilerインストール
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

protoファイル（Procedure定義）からGoコードを生成
```bash
cd api
protoc --go_out=../pkg/grpc --go_opt=paths=source_relative \
	--go-grpc_out=../pkg/grpc --go-grpc_opt=paths=source_relative \
	hello.proto
```


gRPCurlインストール  
※使用するにはサーバーリフレクションが有効になっている必要がある
```bash
curl -LO https://github.com/fullstorydev/grpcurl/releases/download/v1.9.1/grpcurl_1.9.1_linux_x86_64.tar.gz
```
