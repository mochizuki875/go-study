package main

import "fmt"

var i int
var j int = 42
var k = 69 // 推定型

func main(){
	a := 2.0 // 推定型(関数内のみ)
	fmt.Println(i,j,k,a)
}