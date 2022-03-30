package main
import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13} // 配列リテラル
	fmt.Println("q= ",q)

	r := []bool{true, false, true, true, false, true} // 配列リテラル
	fmt.Println("r= ",r)

	// 配列型のstruct
	s := []struct { // i,bをフィールドに持つstruct
		i int
		b bool
	}{
		{2, true}, // sliceを要素として取得
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println("s= ",s)

}