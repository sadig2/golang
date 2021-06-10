package main

import "fmt"

func sum(x, y int, c chan int) {
	sum := x + y
	c <- sum
}

func main() {

	c := make(chan int)

	go sum(2, 3, c)
	go sum(5, 6, c)
	go sum(1, 2, c)

	for i := range c {
		fmt.Println(i)
	}

	a, b, d := <-c, <-c, <-c

	fmt.Println(a, b, d)

}
