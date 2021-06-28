package main

import (
	"bufio"
	"fmt"
	"strings"
)

type something interface {
	bark()
}

type XRen uint64

func (x XRen) bark() {
	fmt.Println("dog is barking")
}

func main() {

	var x something = XRen(4)
	x.bark()

	reader := strings.NewReader("hey  fewf\nv f344 ")
	scanner := bufio.NewScanner(reader)

	i := 0
	fmt.Println(scanner.Text())

	for i < 2 {
		fmt.Println(scanner.Text())
		fmt.Println(i)
		i += 1
	}

}
