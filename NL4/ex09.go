package main

import "fmt"

func main() {
	x := map[string][]string{
		"bond_james":      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	x["powers_austin"] = []string{"shagging", "loving", "baby yeah"}

	for i, v := range x {

		fmt.Printf("This is the record for %v:\n", i)

		for j, val := range v {
			fmt.Printf("\t%v\t %v \n", j, val)
		}
	}

}
