package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Binary<--->Decimal Converter --- Enter a number to conver")
	var input int
	flag := false
	fmt.Scanln(&input)
	str := strconv.Itoa(input)

	for i := 0; i < len(str); i++ {
		if str[i] > 1 {
			flag = true
			break
		}
	}
	if flag {
		convertToBinary(input)
	} else {
		convertToDec(input)
	}
	//in Progress
	//add control to decide which function should be called based on user input
}

func convertToBinary(x int) {
	fmt.Printf("%b\n", x)
}

func convertToDec(x int) {
	fmt.Printf("%d\n", x)
}
