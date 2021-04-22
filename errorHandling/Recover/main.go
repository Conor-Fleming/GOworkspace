package main

import "fmt"

//simple program illustrating how recover wrks with deferred functions

func main() {
	var x int
	x++
	fmt.Println(x)
	i := c()
	fmt.Println(i)
	f() //call to f **step1**
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() { //function is deferred so program flow will move to line 22 **step 2**
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0) //**step3** call to g function
	fmt.Println("Returned normally from g.")
}

func g(i int) { //this function runs recursively eventually incrementing i till the program falls into the if
	//statement on line 28. this will then call panic which triggers all the defered statements/functions back up the call stack
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}
