package main

import "fmt"

//simple pointer dereferencing exercise

type person struct {
	first string
	last  string
	age   int
}

func changeMe(x *person) {
	x.age = 1
	x.first = "Cormac"
	x.last = "Liam"
}

func main() {
	val := 50
	fmt.Println(&val)

	newPers := person{
		first: "Conor",
		last:  "Fleming",
		age:   24,
	}
	fmt.Println(newPers)
	changeMe(&newPers)
	fmt.Println(newPers)
}
