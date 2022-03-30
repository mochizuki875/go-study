package main
import "fmt"

type Vertex struct{
	X, Y int
}

// リテラル
var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 2}
	v3 = Vertex{}
	p = &Vertex{1, 2} // structへのポインタを返す
)

func main() {
	fmt.Println("v1: ", v1)
	fmt.Println("v2: ", v2)
	fmt.Println("v3: ", v3)
	fmt.Println("p: ", p)
	fmt.Println("&p: ", &p)
	fmt.Println("*p: ", *p)	
	fmt.Println("p.X: ", p.X)
	
}