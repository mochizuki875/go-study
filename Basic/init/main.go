package main

import (
	"fmt"

	"example.com/init/foo"
)

func init() {
	fmt.Println("Call init() in main")
}

func main() {
	fmt.Println("Call main() in main")

	f := foo.Foo{1, 2}
	fmt.Println("f = ", f)

}
