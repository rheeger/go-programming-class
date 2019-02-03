package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (a person) speak() {
	fmt.Printf("Hello my name is %v %v and my age is %v\n", a.first, a.last, a.age)
}
func main() {
	p1 := person{
		first: "Robbie",
		last:  "heeger",
		age:   28,
	}

	p1.speak()
}
