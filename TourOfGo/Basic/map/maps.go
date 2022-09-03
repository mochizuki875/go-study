package main

import "fmt"

// struct定義
type Vertex struct {
	Lat, Long float64
}

func main() {

	// mapはKey-Value
	// mapの初期化
	// stringとVertex型structを紐付け
	m := make(map[string]Vertex)
	// var m = map[string]Vertex{} // これでも良い

	// Key（string）に対してValue（Vertex）を紐付け
	m["Foo"] = Vertex{
		10,
		-20,
	}
	m["Bar"] = Vertex{
		30,
		-74,
	}

	fmt.Println(m["Foo"])
	fmt.Println(m["Bar"])
	fmt.Println("len(m)=", len(m))

}
