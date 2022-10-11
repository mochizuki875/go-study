package main

import (
	"fmt"
	"time"
)

// エラー出力の構造体を定義
type MyError struct {
	When time.Time
	What string
}

// fmtパッケージにてerrorインターフェイスが定義されており、その中でError関数が定義されている。
/* =======================
	type error interface {
		Error() string
	}
======================= */

// &MyError型のインターフェイス実装
func (e *MyError) Error() string {
	// MyErrorの出力をカスタマイズ
	// Sprintf: stringとその他の型をstring型としてまとめて扱う
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// サンプル関数
// error型(interface)を返す
func run() error {
	// return &MyError{ // *main.MyError型が返される
	// 	time.Now(),
	// 	"it didn't work.",
	// }

	// 上記はこれと同義
	err := &MyError{time.Now(), "it didn't work."} // Errorは*MyErrorのpointer-Receiverメソッドなので&MyErrorとする
	return err                                     // 関数の戻り値としてinterfceが指定されてるので a(error型) := v(*MyError型)のように*MyError型の変数にインターフェイスを実装しているイメージ
}

func main() {
	// run()を実行し、戻り値をerrとして取得。(エラー処理に該当)
	// errはrun()の戻り値としてerrorインターフェイスのError()メソッドが実装された&MyError型になる
	if err := run(); err != nil {
		fmt.Printf("(%v, %T)\n", err, err)
		fmt.Println(err) // 出力がError()メソッドで定義した形にカスタマイズされる
	}
}
