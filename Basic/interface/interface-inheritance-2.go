/*
GoでInterfaceと埋め込みを実現する
構造体にインターフェイスを埋め込むパターン

https://qiita.com/tenntenn/items/bba69d84a1e0b67c4db8
https://qiita.com/momotaro98/items/4f6e2facc40a3f37c3c3
*/

package main

import (
	"fmt"
	"reflect"
)

type IHuman interface { // インターフェイス定義
	GetName() string
	GetAge() int
}

// 普通の構造体定義とメソッド実装
type Child struct {
	name string
	age  int
}

func (c *Child) GetName() string { // Childに対するインターフェイス実装
	return c.name
}
func (c *Child) GetAge() int { // Childに対するインターフェイス実装
	return c.age
}

type Adult struct {
	IHuman // Adultに対するIHumanインターフェイスの埋め込み
	smoke  bool
}

func main() {
	var human IHuman // インターフェイス用の変数を用意

	tom := &Child{name: "tom", age: 5}
	// fmt.Println("[Child Fild] ", "name: ", tom.name, " age: ", tom.age)
	fmt.Println(*tom)

	// ChildrenをIHumanに実装
	human = tom // 普通のインターフェイスの使い方
	fmt.Println("[IHuman Func] ", "Name: ", human.GetName(), " Age: ", human.GetAge())
	fmt.Println(reflect.TypeOf(human))

	fmt.Println("=========================")

	// ★インターフェイスを実装していないAdultにChildを埋め込む
	// AdultにIHumanインターフェイスを実装してChildを埋め込む(IHuman→&Child)
	// AdultはChildに実装されたメソッドをあたかも自分のメソッドのように呼び出せる
	john := &Adult{&Child{name: "john", age: 30}, true}
	//fmt.Println("[Adult Fild] ", "name: ", john.name, " age: ", john.age, "smoke: ", *john.smoke)
	fmt.Println("[IHuman Func] ", "GetName: ", john.GetName(), " GetAge: ", john.GetAge())
	fmt.Println(*john)
	// fmt.Println(*&john.name) // nameはAdultのフィールドでないのでアクセスできない(あくまでIHUMANが継承されているだけ)
	fmt.Println(*&john.smoke) // smokeはAdultのフィールドなのでアクセスできる

	human = john
	fmt.Println("[IHuman Func] ", "GetName: ", human.GetName(), " GetAge: ", human.GetAge()) // IHumanが継承されているのでメソッド呼び出しができる
	fmt.Println(reflect.TypeOf(human))

	ihuman, ok := human.(IHuman)

	fmt.Println(ihuman, ok)

}
