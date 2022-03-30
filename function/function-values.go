package main
import (
	"fmt"
	"math"
)

// 引数で指定された関数に所定の値を入れた際の戻り値を返す
func compute(fn func(float64, float64) float64) float64{
	return fn(3,4)
}

func main(){
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	// hypot関数の動作
	fmt.Println(hypot(3, 4))

	// hypot関数をcompute関数に渡した時の動作
	fmt.Println(compute(hypot))

	// math.Pow関数をcompute関数に渡した時の動作
	fmt.Println(compute(math.Pow))
}