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

	x := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range x {

		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, v := range v.favcream {
			fmt.Println("\t", i, v)

		}
	}
}
