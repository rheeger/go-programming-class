package main

import "fmt"

var y bool

type burger bool

var x burger

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = true
	fmt.Print(x)
	y = bool(x)
	fmt.Printf("\t%v", y)
}
