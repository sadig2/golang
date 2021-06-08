package main

import "fmt"

func fib() func() int {
	first := 0
	second := 1
	return func() int {
		sum := first + second
		first = second
		second = sum
		return sum
	}
}

func main() {
	f := fib()

	for i := 0; i < 10; i++ {

		fmt.Println(f())
	}

}
