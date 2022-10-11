/*
 Generics
 どのような型も対応できる汎用型の指定
*/

package main

import "fmt"

// 通常の構造体定義
// type Vertex struct {
// 	A *[]int
// 	B []int
// }

// Genericsを使用した構造体定義
type Vertex[T any] struct {
	A *[]T // 自己参照
	B []T
}

// Genericsを使用した構造体定義
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {

	// Genericsでtypeを定義している場合は初期化時にTが何のtypeなのか指定する必要がある(Vertex[int])
	v := Vertex[int]{&[]int{1, 2, 3}, []int{4, 5, 6}}
	fmt.Println(v)

	fmt.Println("================")

	list := &List[int]{&List[int]{&List[int]{nil, 3}, 2}, 1}
	fmt.Println("list = ", list)
	fmt.Println("list.next = ", list.next)
	fmt.Println("list.next.next = ", list.next.next)

}
