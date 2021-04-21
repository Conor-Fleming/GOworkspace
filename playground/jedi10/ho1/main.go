package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	//unblock with func literal
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)

	//unblocking with buffered channel
	x := make(chan int, 1)
	x <- 100
	fmt.Println(<-x)
}
