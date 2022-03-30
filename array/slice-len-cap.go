package main
import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("s: ")
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	fmt.Printf("s[:0]: ")
	printSlice(s)

	// Extend its length.
	s = s[:4]
	fmt.Printf("s[:4]: ")
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	fmt.Printf("s[2:]: ")
	printSlice(s)

	s = s[1:4]
	fmt.Printf("s[1:4]: ")
	printSlice(s)

	var nil_slice []int
	fmt.Printf("nil_slice: ")
	printSlice(nil_slice)

}

// sliceの長さと容量を取得
func printSlice(s []int){
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}