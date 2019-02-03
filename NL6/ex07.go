package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}

	a := func() {
		total := 0
		for _, v := range x {
			total += v
		}
		fmt.Println("The anonymous func gave us this: ", total)

	}

	a()

}
