package main

import "fmt"

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func factorialLoop(x int) int {
	total := 1
	for i := x; i > 0; i-- {
		total *= i
	}
	return total
}

func main() {
	fmt.Println(factorial(3))
	fmt.Println(factorialLoop(3))

}
