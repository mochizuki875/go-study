# Exercise: Slices
A Tour of Goより
https://go-tour-jp.appspot.com/moretypes/18

2次元sliceを渡すとイメージバイナリを返却する。

独自パッケージが依存関係に含まれるため以下手順で取得
https://qiita.com/uchiko/items/64fb3020dd64cf211d4e

~~~
❯❯❯ mkdir exercise-slices

go.modとgo.sumが作成される
❯❯❯ go mod init example.com/exercise-slices
go: creating new go.mod: module example.com/exercise-slices
go: to add module requirements and sums:
        go mod tidy

❯❯❯ go mod tidy
go: finding module for package golang.org/x/tour/pic
go: downloading golang.org/x/tour v0.1.0
go: found golang.org/x/tour/pic in golang.org/x/tour v0.1.0
~~~

## pic module

https://pkg.go.dev/golang.org/x/tour/pic

~~~
func Show(f func(dx, dy int) [][]uint8)
~~~