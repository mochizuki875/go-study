package main
import "fmt"

type Vertex struct {
	X, Y float64
}

// ポインタレシーバー
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f 
}

// ポインタを引数に取る通常の関数
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f 
}

func main() {
	v := Vertex{3,4}
	v.Scale(2) // (&v).Scale()と解釈されている
	ScaleFunc(&v, 10)

	fmt.Println("v:", v)


	p := &Vertex{3,4}
	p.Scale(3) 
	ScaleFunc(p, 10)

	fmt.Println("p:", p) // ポインタを返す
	fmt.Println("&p:", &p) // ポインタアドレスを返す
	fmt.Println("*p:", *p) // 実際の値を返す
	
}