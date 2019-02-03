package main

import "fmt"

type person struct {
	first string
	last  string
}

func changeme(p *person) {
	p.first = "Arnold"

}

func main() {
	p1 := person{
		first: "Robbie",
		last:  "heeger",
	}
	fmt.Println(p1)
	changeme(&p1)
	fmt.Println(p1)

}
