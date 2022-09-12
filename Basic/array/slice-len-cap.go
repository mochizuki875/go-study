package main

import "fmt"

// sliceの長さと容量を取得
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	// sliceを作成
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("s: ", s) // sliceを表示 [0 1 2 3 4 5]
	printSlice(s)         // sliceのlength,capacityをsliceと合わせて表示 (length, capacity)=(6,6)

	fmt.Println("==========================")

	// Slice the slice to give it zero length.
	s = s[:0]
	fmt.Println("s[:0]: ", s)
	printSlice(s) // (length, capacity)=(0,6) capacityは元の配列の長さ分確保される

	fmt.Println("==========================")

	// Extend its length.
	s = s[:4]
	fmt.Println("s[:4]: ", s) // 0~3要素まで [0 1 2 3]
	printSlice(s)             // (length, capacity)=(4,6)

	fmt.Println("==========================")

	// Drop its first two values.
	s = s[2:]                 // 上で作成したsliceからsliceを作成
	fmt.Println("s[2:]: ", s) // 2~最後(3)の要素まで [2 3]
	printSlice(s)             // (length, capacity)=(2,4)

	fmt.Println("==========================")

	s = s[1:4]              // sliceの拡張(re-slice) 上の配列にない要素をsliceしようとしているのではみ出た分は元のsからslice
	fmt.Println("s[1:4]: ") // 1~3要素まで [3 4 5]
	printSlice(s)           // (length, capacity)=(3,3) capacityもresliceに合わせて増えた

	fmt.Println("==========================")

	var nil_slice []int // 初期化していない宣言だけのslice
	fmt.Println("nil_slice: ")
	printSlice(nil_slice) // (length, capacity)=(0,0)

}
