package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	// stringをスペースで区切る
	ss := strings.Fields(s)

	// mapを定義
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
	wc.Test(WordCount)
}