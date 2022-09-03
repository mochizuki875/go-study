package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Stringerインターフェイスはfmtパッケージで定義されている
// 出力をカスタマイズできる
// Stringerインターフェイスで定義されているString()関数を定義することでPersonに実装される

// func (p Person) String() string {
// 	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
// }

func (p *Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	// a := Person{"Tanaka", 20}
	// z := Person{"Goto", 30}

	a := &Person{"Tanaka", 20} // String()はPointer-ReceiverなのでPersonとしてはダメ
	z := &Person{"Goto", 30}

	fmt.Println(a, z) // 出力がString()メソッドで定義した形にカスタマイズされる
}
