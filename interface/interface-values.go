package main
import (
	"fmt"
	// "math"
)

// インターフェイスの定義
type I interface {
	M()
}

// T型の定義
type T struct {
	S string
}

// T型におけるインターフェイス実装
func (t T) M() {
	fmt.Println(t.S)
}

// F型の定義
type F float64

// F型におけるインターフェイス実装
func (f F) M() {
	fmt.Println(f)
}

// インターフェイスの型を表示
// (value, type)
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I // インターフェイス変数の用意

	i = T{S: "Foo"}
	describe(i)
	i.M()

	i = F(-100)
	describe(i)
	i.M()
}