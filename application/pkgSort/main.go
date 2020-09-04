package main

import (
	"fmt"
	"sort"
)

type Person struct { //
	First string
	Age   int
}

type ByAge []Person //setup to sort type Person by age

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []Person //setting up to sort by type name

func (name ByName) Len() int           { return len(name) }
func (name ByName) Swap(i, j int)      { name[i], name[j] = name[j], name[i] }
func (name ByName) Less(i, j int) bool { return name[i].First < name[j].First }

//demonstrating sort function from package sort with slice of strings and ints
func main() {
	xi := []int{24, 6, 75, 98, 12, 4, 2, 96, 33}
	xs := []string{"this", "is", "a", "silly", "stupid", "silly", "fun", "string"}

	fmt.Println(xi)
	fmt.Println(xs)

	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println(xi)
	fmt.Println(xs)

	fmt.Println("*********")
	fmt.Println("*********")

	//custom sort

	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)
}
