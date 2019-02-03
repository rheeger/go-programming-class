package main

import "fmt"

const (
	x = 2019
	y = x - iota
	z = x - iota
	a = x - iota
)

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)
}
