package main

import (
	"fmt"
)

type person struct {
	first string
}

type human interface {
	speak()
}

func (a *person) speak() {
	fmt.Println("AHHHHHHHHH")
}

func saysomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "James",
	}

	saysomething(p1)
	saysomething(&p1)

}
