package main

import "fmt"

func main() {
	a := 42
	fmt.Printf("%v\t%#x\t%b\t\n", a, a, a)
	b := a << 1
	fmt.Printf("%v\t%#x\t%b\t\n", b, b, b)
}
