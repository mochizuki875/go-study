package main
import "fmt"

func echo() string {
	// 遅延処理
	// 呼び出し元の関数がreturnする時に実行
	defer fmt.Println("Echo from Function defer")

	fmt.Println("Echo from Function")
	return "Function done"
}

func main() {
	fmt.Println(echo())	
	fmt.Println("Echo From main")

}