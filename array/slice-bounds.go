package main
import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s= ", s)

	s1 := s[1:4]
	fmt.Println("s[1:4]= ", s1)

	s1 = s[:2]
	fmt.Println("s[:2]= ", s1)

	s1 = s[1:]
	fmt.Println("s[1:]= ", s1)

}