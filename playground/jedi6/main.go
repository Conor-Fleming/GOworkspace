package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Hello, my name is", p.first, p.last, "and i am", p.age, "years old.")
}

func gettingDeffered() {
	fmt.Println("Call me first but run me last")
}

func foo() int {
	return 25
}
func bar() (int, string) {
	return 26, "whats funnier than 25?"
}

func foo2(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func anotherReturned() func() string {
	return func() string {
		return "its me im being returned:)"
	}
}

func main() { //
	defer gettingDeffered() //using defer to show deferred func runs after containing scope exits
	func1 := foo()
	func2, func2_1 := bar()
	fmt.Println(func1, func2, func2_1)

	//using sum function "foo2"
	vals := []int{3, 5, 7, 9, 1}
	result := foo2(vals...)
	fmt.Println(result)

	Conor := person{
		first: "Conor",
		last:  "Fleming",
		age:   24,
	}
	Conor.speak()
	declarations()

	//exercise 6 anon func
	func() {
		fmt.Println("this bitch anonymous boi")
	}()
	//assigning a function to a variable and calling variable
	sw := func() {
		fmt.Println(`this bitch anonymous boi and its assigned to "sw"`)
	}
	sw()

	//exercise 8
	returnedstring := anotherReturned()
	fmt.Println(returnedstring())

	//closure work
	x := 49
	fmt.Println(x)
	{
		x := 50
		fmt.Println(x)
	}
	fmt.Println(x)

}
