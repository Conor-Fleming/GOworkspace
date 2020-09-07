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
	//saySomething(conor)
	saySomething(&conor)

}

func saySomething(h human) {
	h.speak()
}
