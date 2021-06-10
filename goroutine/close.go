package main

import "fmt"

func fibon(n int, c chan int) {

	x, y := 0, 1

	for i := 0; i < n; i++ {
		x, y = y, x+y

		c <- x
	}

	close(c)

}

func main() {

	c := make(chan int, 5)

	fibon(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}

	// var arr []int

	// arr = []int{
	// 	3, 4, 5,
	// }

	// for _, i := range arr {
	// 	fmt.Println(i)
	// }
}
