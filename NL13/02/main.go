package main

import (
	"fmt"
	"github.com/rheeger/go-programming-class/NL13/quote"
	"github.com/rheeger/go-programming-class/NL13/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
