package main

import "fmt"

func main() {
	favsport := "baseball"
	switch favsport {
	case "not favSport":
		fmt.Println("this doesn't print")
	case "baseball":
		fmt.Println("Take me out to the ball game")

	}
}
