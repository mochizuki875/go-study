package main
import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S) // t=nilの場合、不正メモリアクセスとしてSIGSEGVが返されてしまうので上でnilの場合の処理を実装している
}

func describe(i I) {
	fmt.Println("(%v, %T)=", i, i)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()

}