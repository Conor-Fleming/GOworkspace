package main

import (
	"fmt"
	"sort"
)

type userzz struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []userzz //setup to sort type Person by age

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []userzz //setting up to sort by type name

func (name ByName) Len() int           { return len(name) }
func (name ByName) Swap(i, j int)      { name[i], name[j] = name[j], name[i] }
func (name ByName) Less(i, j int) bool { return name[i].Last < name[j].Last }

func sorting() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

	u1 := userzz{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := userzz{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := userzz{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []userzz{u1, u2, u3}

	//fmt.Println(users)

	sort.Sort(ByAge(users))
	fmt.Println(users)
	fmt.Println()
	sort.Sort(ByName(users))
	fmt.Println(users)

	for _, v := range users {
		sort.Strings(v.Sayings)
		fmt.Println(v.Sayings)
		fmt.Println()
	}
}
