package main

import "fmt"

func main() {

	a := outer()
	a()

}

func outer() func() {
	return func() {
		fmt.Println("This is the output of the Func returned from the outer func")
	}
}
