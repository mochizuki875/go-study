/*
arrayとsliceの比較
*/

package main

import "fmt"

func main() {
	primes := [6]int{1, 2, 3, 4, 5, 6} // 配列を定義（配列は固定長）
	fmt.Println("primes= ", primes)

	primes2 := []int{1, 2, 3, 4, 5, 6} // lengthを定義しない動的配列(slice)
	fmt.Println("primes2= ", primes2)

	var s []int = primes[1:4] // sliceを取得（sliceは可変長）要素1以上~4未満
	fmt.Println("primes[1:4]= ", s)

	s[0] = 100 // sliceは配列への参照なので元の配列の値も更新される

	fmt.Println("primes= ", primes)
	fmt.Println("primes[1:4]= ", s)

}
