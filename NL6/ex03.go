package main

import "fmt"

func main() {
	defer fmt.Println("there it is")

	fmt.Println("Whoop")
}
