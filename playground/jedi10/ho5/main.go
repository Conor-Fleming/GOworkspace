package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1) //buffered channel

	c <- 42

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c //using comma ok idiom to check if channel is open/close, should return 0 and false
	fmt.Println(v, ok)
}
