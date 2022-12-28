# Context
タイムアウトやキャンセルなどの情報をcontextと言う単位で扱う。
`goroutine`などスレッドを跨ぐ場合に便利

- 処理の締め切りを伝達
- キャンセル信号の伝播
- リクエストスコープ値の伝達

## Contextの生成
### Background
基本こちらを使う。  
空の`Context`を作成する。  

https://pkg.go.dev/context#Background  
```go
ctx := context.Background()
```

### TODO
空の`Context`を作成する。  
暫定的に`Context`を作りたいときに使う？  

https://pkg.go.dev/context#TODO  
```go
ctx := context.TODO()
```

## キャンセル処理
### WithCancel
`context.WithCancel`メソッドで`Context`を生成すると`Done`チャンネルを持つ`Context`を生成する。  
生成した`Done channel`は戻り値として返されたキャンセルメソッドを実行することでCloseできる。  

https://pkg.go.dev/context#WithCancel  

### WithDeadline
`context.WithDeadline`メソッドで`Context`を生成すると`Done`チャンネルを持つ`Context`を生成する。  
生成した`Done channel`は`Deadline`の時刻を過ぎると勝手にクローズされる。  
※戻り値として返されたキャンセルメソッドを実行することでもCloseできる。  

https://pkg.go.dev/context#WithDeadline  

### WithTimeout
`context.WithTimeout`メソッドで`Context`を生成すると`Done`チャンネルを持つ`Context`を生成する。  
生成した`Done channel`は`Timeout`で設定した時間を経過すると勝手にクローズされる。  
`WithDeadline(parent, time.Now().Add(timeout))`と同義。  
※戻り値として返されたキャンセルメソッドを実行することでもCloseできる。  

https://pkg.go.dev/context#WithTimeout

## リクエストスコープの伝播
### WithValue
`context. WithValue`メソッドで`Context`を生成するとKey-Valueを持つ`Context`を生成する。  
`Context`を渡された`goroutine`から`Context`を介して`Key-Value`を取り出せる。  

https://pkg.go.dev/context#WithValue

---

### 参考
- context(https://pkg.go.dev/context)
- golang contextの使い方とか概念(contextとは)的な話(https://qiita.com/marnie_ms4/items/985d67c4c1b29e11fffc)
- context.Context(https://pkg.go.dev/context#Context)
- contextの概要(https://zenn.dev/hsaki/books/golang-context/viewer/definition)
- Doneメソッド(https://zenn.dev/hsaki/books/golang-context/viewer/done)