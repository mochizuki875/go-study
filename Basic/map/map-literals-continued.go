package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Foo": {
		10,
		120,
	},
	"Bar": {
		-20,
		30,
	},
}

func main() {
	fmt.Println(m)
}
