package main

import "fmt"

type Xtype int64

func (i Xtype) String() string {
	return fmt.Sprintf("%d\n - is number", i)
}

func ins(i int) int {
	return i + 2
}

func fins(f func(i int) int) int {
	x := f(5)
	return x + 10

}

func main() {

	result := fins(ins)
	fmt.Printf("result of fins %d \n", result)

	var x Xtype
	x = 5
	fmt.Println(x)
}
