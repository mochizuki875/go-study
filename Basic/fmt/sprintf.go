package main

import (
	"fmt"
	"strings"
)

func main() {
	str := []string{"World", "Tokyo", "Osaka", "Hokkaido"}
	var messages []string
	for _, v := range str {
		// message += fmt.Sprintf("Hello %s, ", v)
		messages = append(messages, fmt.Sprintf("Hello %s", v))
	}

	fmt.Println(messages)
	fmt.Println(strings.Join(messages, ", "))
}
