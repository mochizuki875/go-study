package main
import "fmt"

func main() {
	// makeを使ってゼロ化された動的配列を作る
	a := make([]int, 5) // len=5の配列を作成
	printSlice("a", a)

	var x [5]int
	fmt.Println("x [5]int, ", "len=", len(x), "cap=", cap(x), x)

	b := make([]int, 0, 5) // len=0, cap=5の配列
	printSlice("b", b)

	c := b[:2]
	printSlice("c := b[:2]", c)

	d := b[2:5]
	printSlice("d := b[2:5]", d)

}

func printSlice(s string, x []int){
	fmt.Printf("%s, len=%d, cap=%d %v\n", s, len(x), cap(x), x)
}