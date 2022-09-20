# go study
Go言語学習時のサンプルプログラム集

## 学習教材
### 基本文法
基本文法は以下に沿って学習

- A Tour of Go (https://go-tour-jp.appspot.com/list)

### 配列とSlice
- 構造体を要素に持つスライス (https://nishinatoshiharu.com/initialize-structure-slice/)

### インターフェイス
- Go言語 - 空インターフェースと型アサーション (https://blog.y-yuki.net/entry/2017/05/08/000000)
- Stringerで出力形式をカスタマイズしよう!(https://selfnote.work/20200716/programming/stringer-with-golang/)

### goroutine & channel
goroutineとchannelは頻出かつ難易度が高いので以下コンテンツを利用

- Goで並行処理(基本編)（https://zenn.dev/hsaki/books/golang-concurrency/viewer/intro）
- Goのgoroutine, channelをちょっと攻略！（https://qiita.com/taigamikami/items/fc798cdd6a4eaf9a7d5e）

### Context
タイムアウトやキャンセルなどの情報をcontextと言う単位で扱う。
goroutineなどスレッドを跨ぐ場合に便利

- context(https://pkg.go.dev/context)
- contextの概要(https://zenn.dev/hsaki/books/golang-context/viewer/definition)
- Doneメソッド(https://zenn.dev/hsaki/books/golang-context/viewer/done)

`context.Context型の定義`
```
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```

### httpサーバー
goでhttpサーバーを起動するための`net/http`パッケージの利用方法を以下コンテンツにて学習

- Golang初心者が`net/http`パッケージでWebサーバーをホストする流れを追う（ https://zenn.dev/skonb/articles/0bad1d59371d09 )
- Deep Dive into The Go's Web Server（ https://zenn.dev/hsaki/books/golang-httpserver-internal/viewer/intro ）

### sample-controller
`Kubernetes`の`Controller`に関するサンプル実装
https://github.com/kubernetes/sample-controller/tree/master

## 学習環境

```
❯❯❯ sw_vers
ProductName:    macOS
ProductVersion: 12.3
BuildVersion:   21E230

❯❯❯ go version
go version go1.19 darwin/amd64
```

## Build & Execute

```
go run hello.go
```

## Compile & Execute

```
go build hello.go
./hello
```

## 参考
- プログラミング言語Go完全入門(https://docs.google.com/presentation/d/1RVx8oeIMAWxbB7ZP2IcgZXnbZokjCmTUca-AbIpORGk/edit#slide=id.g4f417182ce_0_0)
- WikiBook Go (https://ja.wikibooks.org/wiki/Go)
- とほほのGo言語入門 (https://www.tohoho-web.com/ex/golang.html)
- 【Go】パッケージ/モジュールやgo modコマンドについてまとめ (https://blog.framinal.life/entry/2021/04/11/013819)
- kubecontorller-book-sample-snippet(https://github.com/govargo/kubecontorller-book-sample-snippet)
- Kubernetesのコードリーディングをする上で知っておくと良さそうなこと(https://bells17.medium.com/things-you-should-know-about-reading-kubernetes-codes-933b0ee6181d)
- Kubernetesコントリビューターチートシート(https://github.com/kubernetes/community/blob/master/contributors/guide/contributor-cheatsheet/README-ja.md)
- client-go under the hood
(https://github.com/kubernetes/sample-controller/blob/release-1.25/docs/controller-client-go.md)