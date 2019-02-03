package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		fmt.Printf("\n\n%v", i)
		for j := 65; j <= 90; j++ {
			fmt.Printf("\n\t%#U", j)
		}
	}
}
