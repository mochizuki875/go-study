package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	var s []int   // slice定義
	printSlice(s) // (len,cap) = (0,0)

	// sliceへ新しい要素を追加
	// func append(s []T, vs ...T) []T

	s = append(s, 0) // s = [0]
	printSlice(s)    // (len,cap) = (1,1)

	s = append(s, 1) // s = [0 1]
	printSlice(s)    // (len,cap) = (2,2)

	s = append(s, 2, 3) // s = [0 1 2 3]
	printSlice(s)       // (len,cap) = (4,4)

	s = append(s, 4, 5, 6) // s = [0 1 2 3 4 5 6]
	printSlice(s)          // (len,cap) = (7,8) 追加する要素数がある程度大きい(sのbacking arrayより大きい)とcapを追加で確保？？(ここの規則はそんなに気にしなくてよさそう)

	s = append(s, 7, 8, 9) // s = [0 1 2 3 4 5 6 7 8 9]
	printSlice(s)          // (len,cap) = (10,16) 追加する要素数がある程度大きいとcapを追加で確保？？

}
