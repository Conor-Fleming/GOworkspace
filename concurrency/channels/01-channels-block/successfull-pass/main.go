package main

import "fmt"

func main() {
	//successful pass example 1
	c := make(chan int)
	go func() {
		c <- 50
	}()

	fmt.Println(<-c)
}
