package main

import "fmt"

func main() {
	for x := 10; x <= 100; x++ {
		if x%2 == 0 {
			fmt.Printf("%v\n", x)
		}
	}
}
