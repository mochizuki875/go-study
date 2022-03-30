package main
import (
	"fmt"
	"math"
)

const Pi = 3.14

func main(){
	const World = "世界"
	fmt.Println("Hell", World)
	fmt.Println("Pi: ", Pi)
	fmt.Println("math.Pi: ", math.Pi)  // mathライブラリでexportされた変数を使用

	const Truth = true
	fmt.Println("Go rules?", Truth)
	
}