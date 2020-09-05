package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user2 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func encoding() {
	u1 := user2{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user2{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	err := json.NewEncoder(os.Stdout).Encode(u2)
	if err != nil {
		fmt.Println("something went wrong")
	}
	fmt.Println()
	err = json.NewEncoder(os.Stdout).Encode(u1)
	if err != nil {
		fmt.Println("something went wrong")
	}

}
