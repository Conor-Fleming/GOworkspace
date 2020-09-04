package main

import "fmt"

func main() {
	x := 0
	fmt.Println(&x)
	fmt.Println(x)
	foo(&x)
	fmt.Println(&x)
	fmt.Println(x)
}

func foo(y *int) {
	fmt.Println(y)
	fmt.Println(*y)
	*y = 43
	fmt.Println(y)
	fmt.Println(*y)
}

//Cleaning my workspace

/*
type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() { //attaching speak function to secretAgent struct *Method*
	fmt.Println("I am", s.first, s.last)
}

func (p person) speak() { //attaching speak function to secretAgent struct *Method*
	fmt.Println("I am", p.first, p.last, "- the person speak")
}

type human interface { //"anything with method "speak()" is also of type human"
	speak()
}

func bar(h human) {
	fmt.Println("I was passed into bar", h)
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	saCon := secretAgent{
		person: person{
			first: "Conor",
			last:  "Fleming",
		},
		ltk: true,
	}

	p1 := person{
		first: "Doctor",
		last:  "Yes",
	}
	fmt.Println(sa1)
	sa1.speak()
	saCon.speak()

	fmt.Println(p1)

	bar(sa1)
	bar(saCon)
	bar(p1)
	testa := func() {
		fmt.Println("you can assign funcitons to a variable.")
	}
	testa()

	const (
		_  = iota
		kb = 1 << (iota * 10)
		mb = 1 << (iota * 10)
		gb = 1 << (iota * 10)
	)

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)
	xi := []int{3, 5, 7, 8, 24, 56, 79}

	foo(1, 2, 3, 4, 5, 6, 7, 8)
	foo(xi...) //"unfurling" a slice of ints to use for a variadic parameter int "...int"
}

//veriadic param
func foo(x ...int) { // zero or more of type int
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position, ", i, "we are now adding,", v, "to the total, which is now ", sum)
	}
	fmt.Println("The total is,", sum)
}

func loopFactorial(fact int) int {
	total := 1
	for x := fact; x > 0; x-- {
		total *= x
	}
	return total
}*/
