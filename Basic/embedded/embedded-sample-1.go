package main

import "fmt"

type str_a struct {
	X int
	Y int
}

type str_b struct {
	*str_a
	Z int
}

func main() {
	a := &str_a{X: 1, Y: 2}
	b := str_b{a, 3}
	b.Z = 3

	fmt.Printf("b.a: %+v", b.a)
}
