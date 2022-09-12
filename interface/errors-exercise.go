package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// 	return fmt.Sprint("cannot Sqrt negative number: %v", e) // これだとエラー %vでErrNegativeSqrt型を扱えないから？
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {

	// エラー処理
	if x < 0 {
		err := ErrNegativeSqrt(x)
		return x, err
	}

	z := float64(1)

	// 平方根算出(整数(ニュートン法))
	for z*z <= x {
		z++
	}
	return z - 1, nil
}

func main() {
	fmt.Println(Sqrt(9))
	fmt.Println(Sqrt(-2))
}
