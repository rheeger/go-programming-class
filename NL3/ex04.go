package main

import "fmt"

func main() {
	y := 1990
	for {
		if y >= 2019 {
			break

		}
		fmt.Println(y)
		y++
	}
}
