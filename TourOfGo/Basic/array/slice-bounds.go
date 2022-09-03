package main

import "fmt"

func main() {
	// sliceのリテラル(初期化)
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s= ", s)

	// 1要素目~3要素目まで
	s1 := s[1:4]
	fmt.Println("s[1:4]= ", s1)

	// 0要素目~1要素目まで
	s1 = s[:2]
	fmt.Println("s[:2]= ", s1)

	// 1要素目~最後の要素目まで
	s1 = s[1:]
	fmt.Println("s[1:]= ", s1)

}
