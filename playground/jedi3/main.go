package main

import "fmt"

func main() {
	//excersise 1 print numbers 1 to 10,000
	//*****************************
	/*for i := 0; i <= 10000; i++ {
	//	fmt.Print(i, " ")
	}*/

	//excersise 2 Print every rune code point of the uppercase alphabet three times.
	//*************************
	/*for i := 65; i < 91; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%#U\n", i)
		}
	}*/

	//excersise 3 for loop

	/*for i := 2020; i > 1994; i-- {
		fmt.Print(i, " ")
	}*/

	//excersise 4 for loop as while
	/*x := 1995
	for {
		if x < 2021 {
			fmt.Print(x, " ")
			x++
		} else {
			break
		}
	}*/

	//excersise 5 Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.
	//*******
	/*for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}*/

	//Ex 6 if statement
	/*test := 26
	if test == 26 {
		fmt.Println("true")
	}*/

	//ex 8
	switch {
	case (7 == 4):
		fmt.Println("y")
	case (5 == 5):
		fmt.Println("wwwww")
	case true:
		fmt.Println("test")
	case (1 == 2):
		fmt.Println("working?")
	}

}
