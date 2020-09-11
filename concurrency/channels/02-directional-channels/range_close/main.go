package main

import "fmt"

//using range to pull values from channel until channel is closed

func main() {
	c := make(chan int)
	go foo(c)

	for v := range c {
		fmt.Println(v)
	}
	//bar(c)
	fmt.Println("about to exit")
}

//send
func foo(c chan<- int) {
	for i := 0; i < 20; i++ {
		c <- i
	}
	close(c)
}
