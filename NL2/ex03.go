package main

import "fmt"

const (
	x         = 42
	y float64 = 42.104
)

func main() {
	fmt.Println(x, y)
	fmt.Printf("%T\t%T\t", x, y)

}
