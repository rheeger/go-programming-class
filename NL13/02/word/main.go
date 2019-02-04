// Package Word returns fun info about the words in Package Quote
package word

import "strings"

// UseCount returns a map with the number of uses of a given word
// with the word next to the number of uses
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}


// Count returns the total number of words in the quote.
func Count(s string) int {
	xs := strings.Fields(s)
	x := 0
	for i := range xs {
		x = i + 1
	}

	return x


}
