/*
  intとmainが呼ばれる順番のチェック
  ①読み込んだパッケージ(foo)のinit
  ②mainのinit()
  ③main()
*/

package foo

import "fmt"

type Foo struct {
	A int
	B int
}

func init() {
	fmt.Println("Call init() in foo")
}

func (f *Foo) sum() int {
	return f.A + f.B
}
