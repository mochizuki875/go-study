/*
関数の引数としてInterfaceを指定し、関数呼び出しに伴う引数代入時に実装するパターン
*/

package main

import "fmt"

type Foo struct {
	a int
	b int
}

func (f *Foo) add() int {
	return f.a + f.b
}

func (f *Foo) mult() int {
	return f.a * f.b
}

type Calc interface{
	add() int
	mult() int
}

func result(c Calc) {
	fmt.Println("add: ",c.add())
	fmt.Println("mult: ",c.mult())

	f,_ := c.(*Foo)
	fmt.Println("f: ",f)
	
}

func main() {

	foo := &Foo{
		a:1,
		b:2,
	}

	result(foo)
}