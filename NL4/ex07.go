package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not Stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	xy := [][]string{x, y}

	for i, v := range xy {
		fmt.Printf("%v\n", i)
		for j, val := range v {
			fmt.Printf("\t index position: %v\t value:\t %v \n", j, val)
		}

	}

}
