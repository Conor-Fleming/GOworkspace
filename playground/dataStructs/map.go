package main

import "fmt"

var m = map[string]int{
	"Conor":  5144505,
	"cormac": 5146445,
}

func introMap() {
	fmt.Println(m["Conor"])
	fmt.Println(m["cormac"])
	if v, ok := m["Conor"]; ok {
		fmt.Println("this is the value check print statement", v)
	}
}

func addMapVal() {
	m["mom"] = 5142529
	m["dad"] = 5148380
	if v, ok := m["mom"]; ok {
		fmt.Println("this is the value check print statement", v)
	}
}

func mapRange() {
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func mapDelete() {
	delete(m, "Conor")
	fmt.Println(m)
}
