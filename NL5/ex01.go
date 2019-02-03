package main

import "fmt"

type person struct {
	first    string
	last     string
	favcream []string
}

func main() {
	p1 := person{
		first: "Robbie",
		last:  "Heeger",
		favcream: []string{"Mint Chip",
			"chocolate",
			"vanilla"},
	}

	p2 := person{
		first: "Ariel",
		last:  "Friedman",
		favcream: []string{"strawberry",
			"vanilla",
			"rocky road"},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favcream {
		fmt.Println("\t", i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favcream {
		fmt.Println("\t", i, v)
	}
}
