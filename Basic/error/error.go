package main

import (
	"errors"
	"fmt"
)

func main() {
	e := errors.New("This is an error message")

	err := fmt.Errorf("ERROR: %s", e)
	fmt.Printf("%v\n", err)

	err_wrapped := fmt.Errorf("ERROR: %w", e)      // %wを使うとエラーをUnwrapするためのUnwrap()が実装される
	fmt.Printf("%v\n", err_wrapped)                // Errorfでerrに設定したメッセージ（WRAPされたもの）を返す
	fmt.Printf("%v\n", errors.Unwrap(err_wrapped)) // eのメッセージだけを返す

}
