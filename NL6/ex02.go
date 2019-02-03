package main

import "fmt"

func main() {
	a := []int{42, 43, 44, 45, 46, 47}
	n := foo(a...)
	fmt.Println(n)

	b := []int{1, 2, 3, 4, 5, 6, 7}
	n2 := bar(b)
	fmt.Println(n2)
}

func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func bar(zi []int) int {
	total := 0
	for _, v := range zi {
		total += v
	}
	return total
}
