package main

import "fmt"

func main() {
	x:= loopFact(4)
	fmt.Println(x)
}


func loopFact(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
