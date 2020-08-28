package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func structInit() {
	p1 := person{
		first: "Conor",
		last:  "Fleming",
		age:   24,
	}

	p2 := person{
		first: "Ciara",
		last:  "Kyne",
		age:   23,
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func embedStruct() { //embedding structs *POLYMORPHISM* *INHERITENCE*

	sa1 := secretAgent{
		person: person{
			first: "Con",
			last:  "Flem",
			age:   100,
		},
		ltk: true,
	}
	fmt.Println(sa1)
	fmt.Println(sa1.first)
	fmt.Println(sa1.person)
	fmt.Println(sa1.age)
	fmt.Println(sa1.person.age) //same as above, dont need to specify parent class unless there is a naming clash

}
