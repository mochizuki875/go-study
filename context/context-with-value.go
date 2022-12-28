package main

import (
	"context"
	"fmt"
)

func main() {

	// 第一引数のContextが第二引数で渡されたものに合致したらValueを表示
	f := func(ctx context.Context, k string) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value: ", v)
			return
		}
		fmt.Println("key not found: ", k)
	}

	// Key-Valueとして"language:Go"を持つContextを生成
	k := "language"
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, "color")
}
