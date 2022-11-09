package main

import (
	"fmt"
	"reflect"
)

func main() {
	const upper = 'A'
	const lower = 'a'
	const toLower = (lower - upper)
	fmt.Println("upper=", upper, " : ", reflect.TypeOf(upper))
	fmt.Println("lower=", lower, " : ", reflect.TypeOf(lower))
	fmt.Println("toLower=", toLower, " : ", reflect.TypeOf(toLower))

	fmt.Println("=====================")

	desc := string("-Aa")
	for i := range desc {
		fmt.Println("i=", i, " : ", reflect.TypeOf(i))
		nm := string(lower + byte(i))
		fmt.Println("nm=", nm, " : ", reflect.TypeOf(nm))
		fmt.Println("=====================")
	}
}
