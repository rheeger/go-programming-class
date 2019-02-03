package main

import "fmt"

var x = 42
var y = "James Bond"
var z = "True"

func main() {
	s := fmt.Sprintf("%v\n%v\n%v\n", x, y, z)
	fmt.Println(s)
}
