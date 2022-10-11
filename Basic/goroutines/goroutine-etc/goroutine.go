package main
import(
	"fmt"
	"time"
)

func f(value string) {
	for i := 0; i < 3; i++ {
		fmt.Println(value)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	go f("Call from goroutine")
	f("Call from main")
	fmt.Println("Done")
}