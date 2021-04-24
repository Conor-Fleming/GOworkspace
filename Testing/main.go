package main

import "fmt"

func main() {
	fmt.Println("4 + 3 =", mySum(4, 3))
	fmt.Println("5 + 5 =", mySum(5, 5))
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
