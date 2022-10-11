package main
import(
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)
	
	go func(){
		time.Sleep(5*time.Second)
		c2 <- "two"
	}()

	go func() {
		time.Sleep(1*time.Second)
		c1 <- "one"
	}()

	for i := 0; i < 2; i++{
		select{
		case msg1 := <- c1: // c1から値を受信した場合
			fmt.Println("Received", msg1)
		case msg2 := <- c2: // c2から値を受信した場合
			fmt.Println("Received", msg2)
		}
	}
}