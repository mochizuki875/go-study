package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// Pointer Receiverは直接参照なのでポインターを引数にとるケースと同じ

// Value Receiver
func (v Vertex) AddFunc() {
	v.X += v.X
	v.Y += v.Y
}

// Pointer Receiver
func (v *Vertex) Add() {
	v.X += v.X
	v.Y += v.Y
}

// Function pointer arg
func AddPointer(v *Vertex) {
	v.X += v.X
	v.Y += v.Y
}

func main() {

	fmt.Println("Value Receiver")
	v1 := Vertex{2, 3}
	fmt.Println(v1)
	v1.AddFunc()
	fmt.Println(v1) // {2 3}

	fmt.Println("============")

	fmt.Println("Pointer Receiver")
	v2 := Vertex{2, 3}
	fmt.Println(v2)
	v2.Add()
	fmt.Println(v2) // {4 6}

	p := &v2
	p.Add() // これでも良い
	fmt.Println(v2)

	fmt.Println("============")

	fmt.Println("Function pointer arg")
	v3 := Vertex{2, 3}
	fmt.Println(v3)
	AddPointer(&v3)
	fmt.Println(v3) // {4 6}

}
