package main
import "fmt"

func adder() func(int) int{ // 戻り値にint型を返す関数を指定
	sum := 0
	return func (x int) int { //戻り値として無名関数（クロージャー）を返す
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(2*i),
		)
	}
}