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

### Queue
- `slice`や`channel`を用いて実現できる
- workqueue(https://qiita.com/tatsuhiro-t/items/49043fca96e484de6261)
  - `Kubernetes`の提供する`client-go`が持つ`WorkQueue`の仕組み

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

### Cobra
コマンドラインインターフェイスを作成するためのライブラリ

- spf13/cobra(https://github.com/spf13/cobra)
- Cobra Generator コマンドテンプレートの作成(https://github.com/spf13/cobra-cli/blob/main/README.md)
- User Guide コマンドの実装例(https://github.com/spf13/cobra/blob/main/user_guide.md)
- Go初心者がcobraを使ってコマンドラインツールを作ってみた話(https://blog.engineer.adways.net/entry/advent_calendar_2018/18)
- 【Golang】cobraで作ったコマンドラインツール(CLI)にフラグを追加する (pflag)(https://simple-minds-think-alike.moritamorie.com/entry/add-flags-to-cobra-app)

### httpサーバー
goでhttpサーバーを起動するための`net/http`パッケージの利用方法を以下コンテンツにて学習  

- Golang初心者が`net/http`パッケージでWebサーバーをホストする流れを追う（ https://zenn.dev/skonb/articles/0bad1d59371d09 )
- Goのhttp.Handlerやhttp.HandlerFuncをちゃんと理解する(https://journal.lampetty.net/entry/understanding-http-handler-in-go)
- Deep Dive into The Go's Web Server（ https://zenn.dev/hsaki/books/golang-httpserver-internal/viewer/intro ）

### Go Test
- Go言語でテスト go test(https://zenn.dev/ichi320/articles/70efa1f122cf8c)

```
$ go test
$ go test -v
$ go test ${DIR}
```

### Kubernetes
- Workqueue Example(https://github.com/kubernetes/client-go/tree/master/examples/workqueue)
- sample-controller `Kubernetes`の`Controller`に関するサンプル実装(https://github.com/kubernetes/sample-controller/tree/master)

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