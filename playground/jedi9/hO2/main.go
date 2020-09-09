package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	first string
	last  string
}

func (p *person) speak() {
	fmt.Println("Hello, i am", p.first, p.last, ".")
}

func main() {
	conor := person{
		first: "Conor",
		last:  "Fleming",
	}
	//saySomething(conor) //demonstrating that you cannot pass in type person if you implement the speak method with a pointer reciever
	saySomething(&conor) //must pass in the value as a pointer

	//however you could call speak straight from type person and bypass the interface
	conor.speak()
}

func saySomething(h human) {
	h.speak()
}
