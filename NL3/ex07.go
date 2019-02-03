package main

import "fmt"

func main() {
	for x := 10; x <= 100; x++ {
		if x%2 == 0 {
			fmt.Printf("%v is divisible by 2\n", x)
		} else if x%3 == 0 {
			fmt.Printf("%v is divisible by 3\n", x)

		} else {
			fmt.Printf("%v is not divisible by 2 or 3\n", x)
		}
	}
}
