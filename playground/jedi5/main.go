package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	favFlavs  []string
}

type vehicle struct { //added for exercise 3, embedding structs
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	//Exercise 1
	//creating two values of type person (p1, p2)
	p1 := person{
		firstName: "Conor",
		lastName:  "Fleming",
		favFlavs:  []string{"vanilla", "Cookie Dough", "Sorbet"},
	}

	p2 := person{
		firstName: "Ciara",
		lastName:  "Kyne",
		favFlavs:  []string{"chocolate", "strawberry", "vanilla"},
	}

	fmt.Println(p1) //printing values then using range to print out each of the values in the favorite ice cream flavors slice
	for _, v := range p1.favFlavs {
		fmt.Println(v)
	}

	fmt.Println(p2)
	for _, v := range p2.favFlavs {
		fmt.Println(v)
	}

	//Exercise2
	m := map[string]person{
		"Fleming": p1,
		"Kyne":    p2,
	}

	for k, v := range m {
		fmt.Println(k, v.favFlavs)
	}

	//Exercise3

	truck1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		fourWheel: true,
	}

	sedan1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: false,
	}
	fmt.Println(truck1, "\t", sedan1)

	fmt.Println(truck1.color)
	fmt.Println(sedan1.luxury)

	//Exercise4
	computer := struct {
		size   int
		noisey bool
		color  string
	}{
		size:   30,
		noisey: false,
		color:  "Space Grey",
	}

	fmt.Println(computer)
	fmt.Println(computer.noisey)

}
