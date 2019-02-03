package main

import "fmt"

func main() {
	for i := 0; i < 200; i++ {
		fmt.Printf("%d\t%#x\t%#U\n", i, i, i)
	}
}
