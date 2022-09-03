package main

import (
	"fmt"
	"strings"
	"reflect"

)

func main() {
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))

	// 文字列をスペースで区切る
	s := strings.Fields("  foo bar  baz   ")

	// 型確認
	fmt.Println(reflect.TypeOf(s))

	for index, value :=  range s {
		fmt.Printf("index: %d, value: %s\n", index, value)
	}
}