package main

import "fmt"

func main() {
	for x := 10; x <= 100; x++ {
		fmt.Printf("%v divided by 4 has a remainder of %v\n", x, x%4)
	}
}
