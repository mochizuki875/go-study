package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i                  //point to i (*i型)
	fmt.Println("*p = ", *p) // *pはポインタpの指す値を示す
	fmt.Println("p = ", p)   // pはポインタの値（メモリアドレス）そのものを示す

	*p = 21 // pの示す先（i）に値を格納
	fmt.Println("i(*p) = ", i)

	p = &j // pの示す先をjのアドレスに変更
	*p = *p / 37
	fmt.Println("j = ", j)
	fmt.Println("&j = ", &j)
	fmt.Println("p = ", p)
}
