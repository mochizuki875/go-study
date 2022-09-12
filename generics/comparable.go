package main

import "fmt"

// 引数の汎用化 + 引数型の制約

// 第1引数の配列をチェックし、第2引数に合致する要素が含まれていたらその要素番号を返すIndexと言う関数
// [T comparable](s []T, x T)と言う引数の宣言によりtypeを汎用化している
// [T comparable]によりTの制約を行っている(==や!=で比較可能な型のみTとして許可)
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {

	si := []int{1, 2, 3, 4, -1}
	fmt.Println(si)
	fmt.Println(Index(si, 4))

	fmt.Println("================")

	fmt.Println(si)
	fmt.Println(Index(si, 10))

	fmt.Println("================")

	ss := []string{"aaa", "bbb", "ccc"}
	fmt.Println(ss)
	fmt.Println(Index(ss, "bbb"))

}
