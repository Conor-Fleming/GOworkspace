package main

import (
	"fmt"
	"log"

	"github.com/Conor-Fleming/GOworkspace/documentation/DogPackage"
)

func main() {
	fmt.Println("Enter an age you would like converted to Dog years:")
	var age int
	_, err := fmt.Scanln(&age)
	if err != nil {
		log.Fatalln("That is not a number. Exiting.")
	}

	result := DogPackage.Years(age)
	fmt.Println(age, "in dog years is:", result)
}
