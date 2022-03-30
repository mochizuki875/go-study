# Exercise: Slices
A Tour of Goより
https://go-tour-jp.appspot.com/moretypes/23

渡された文字列内の単語出現回数をカウントするテストスイートを成功させる。
https://github.com/golang/tour/blob/master/wc/wc.go

~~~
❯❯❯ mkdir exercise-maps

❯❯❯ go mod init example.com/exercise-maps
go: creating new go.mod: module example.com/exercise-maps
go: to add module requirements and sums:
        go mod tidy

❯❯❯ go mod tidy
go: finding module for package golang.org/x/tour/wc
go: found golang.org/x/tour/wc in golang.org/x/tour v0.1.0
~~~

## strings.Fields
与えられた文字列をスペースで区切ってstring型の配列を返す。

https://pkg.go.dev/strings#Fields

~~~
func Fields(s string) []string
~~~