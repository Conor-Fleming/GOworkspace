package main

import "fmt"

func main() {
	//c := make(chan int, 2)

	//c <- 1
	//c <- 2
	//fmt.Println(<-c, <-c)
	fmt.Println("----------")
	//fmt.Printf("%T\n", c)

	//directional channels
	c := make(chan int)
	cr := make(<-chan int) //recieve
	cs := make(chan<- int) //send

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

	go func() {
		cs <- 42
	}()

	//fmt.Println(<-cs) //spits an error, send only channel

}
