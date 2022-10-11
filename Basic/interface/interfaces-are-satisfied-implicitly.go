package main
import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// 明示的にIというインターフェイスに実装する宣言をしなくてもIに含まれるメソッドを実装することで自動的に実装される
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"} // T型の構造体をI型のインターフェイスに変換
	i.M() 
}