package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {

	// stringをスペースで区切る
	ss := strings.Fields(s)

	m := make(map[string]int)	

	for i := range ss {
		_, ok := m[ss[i]]
		if ok {
			m[ss[i]] ++
		}  else {
			m[ss[i]] = 1
		}
	}

	return m
}

func main() {
	m := WordCount("A B C A C")
	fmt.Println(m)
}