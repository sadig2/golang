package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := "over\nthe \nrainbow"

	scanner := bufio.NewScanner(strings.NewReader(s))

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
