package main
import "fmt"

func main(){
	primes := [6]int{2,3,5,7,11,13} // 配列を定義（配列は固定長）
	fmt.Println("primes= ", primes)
	var s []int = primes[1:4] // sliceを取得（sliceは可変長）要素1以上~4未満
	fmt.Println("primes[1:4]= ", s)

	s[0] = 100

	fmt.Println("primes= ", primes)
	fmt.Println("primes[1:4]= ", s)


}