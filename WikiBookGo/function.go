package main

import(
	"fmt"
)

func main(){
	fmt.Println(subr(3,2))
}

func subr(x, y int)(int, int){
	return x+y, x-y
}