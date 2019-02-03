package main

import "fmt"

func main() {
	x := [5]int{5, 6, 7, 8, 9}

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", x)

}
