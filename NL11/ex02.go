package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatal("This didn't make it to the JSON parser:", err)
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		err = fmt.Errorf("you haven't passed any data with your byte: %v", bs)
		return []byte{}, err
	}
	return bs, nilw
}
