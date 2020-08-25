package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number to convert it to Binary")
	var input int
	fmt.Scanln(&input)

	convertToBinary(input)
	//in Progress
	//add control to decide which function should be called based on user input
}

func convertToBinary(x int) {
	fmt.Printf("%b\n", x)
}

func convertToDec(x int) {
	fmt.Printf("%d\n", x)
}
