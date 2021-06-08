package main

import (
	"fmt"
	"strings"
)

func wordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range words {
		m[v] = m[v] + 1
	}
	return m
}

func main() {
	word := "privet kak dela dela"
	result := wordCount(word)
	fmt.Println(result)
}
