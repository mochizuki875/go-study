/*
GoでInterfaceと埋め込みを実現する

https://qiita.com/Yuuki557/items/b74d0304d01a90f90947
*/

package main

import "fmt"

type IHuman interface {
	GetName() string
}

type Human struct {
	name string
}

func (h *Human) GetName() string {
	return h.name
}

type Citizen struct {
	Human    // 埋め込み
	language string
}

func (c *Citizen) GetLang() string {
	return c.language
}

func (c *Citizen) ChangeLang(lang string) {
	c.language = lang
}

func main() {
	taro := Citizen{
		Human: Human{
			name: "taro",
		},
		language: "Japanese",
	}

	fmt.Println(taro.GetName())
	fmt.Println(taro.GetLang())

	taro.ChangeLang("English")
	fmt.Println(taro.GetLang())

}
