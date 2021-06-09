package main

import "fmt"

type Merror struct {
	what string
}

func (er *Merror) Error() string {
	return fmt.Sprintf(" %v", er.what)
}

func runError() {
	err := &Merror{
		what: "some fucking error",
	}
	fmt.Println(err)
}

func main() {

	runError()

}
