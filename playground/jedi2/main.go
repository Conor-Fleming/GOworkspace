package main

import "fmt"

const (
	con     = "Conor"
	b   int = 5
)

const (
	a = 2017 + iota
	a2
	a3
	a4
)

func main() {
	x := 3
	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%O\n", x)
	fmt.Println()

	fmt.Println(con, b)
	fmt.Println()
	foo := 12
	fmt.Printf("%d\n", foo)
	fmt.Printf("%b\n", foo)
	fmt.Printf("%#x\n", foo)
	fmt.Println()

	foo2 := foo << 1
	fmt.Printf("%d\n", foo2)
	fmt.Printf("%b\n", foo2)
	fmt.Printf("%#x\n", foo2)

	lit := `this is a raw lit "string" lol`
	fmt.Println(lit)

	fmt.Println(a, a2, a3, a4)

}
