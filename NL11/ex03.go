package main

import (
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

type customErr struct {
}

func (c customErr) Error() string {
	return "This is our new custom Error"
}

func main() {
	c1 := customErr{}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)

}
