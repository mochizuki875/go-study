// ポインタ型と参照型

package main

import "fmt"

type A struct {
	x int
	y int
}

// ポインタが引数になってるので関数内での変更が実態に反映される
func mod(a *A) {
	 a.x = 10
	 a.y = 20
}

// 参照型が引数になっているので関数内で変更を行っても実態は変更されない
func Mod(a A) {
	a.x = 100
	a.y = 200
}

func main(){
	a := &A{x:1,y:2}
	mod(a) // 関数で構造体の中身を書き換え
	fmt.Println(*a)

	fmt.Println("=================")

	Mod(*a) // 関数で構造体の中身を書き換え
	fmt.Println(*a)

}