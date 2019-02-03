package main

import "fmt"

func main() {
	y := foo()
	fmt.Println(y)
}

func foo() int {
	x := 42
	return x
}
