package main

import "fmt"

func main() {
	//successful pass example 1
	c := make(chan int)
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	//successfull buffer example 2
	//be weary with buffer channels
	e := make(chan int, 2)

	e <- 52
	e <- 55
	fmt.Println(<-e, <-e)

}
