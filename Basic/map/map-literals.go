package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// map literal(mapの初期化)
var m = map[string]Vertex{
	"Foo": Vertex{
		10,
		120,
	},
	"Bar": Vertex{
		-20,
		30,
	},
}

func main() {
	fmt.Println(m)
}
