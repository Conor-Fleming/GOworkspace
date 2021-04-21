package main

import "fmt"

func main() {
	c := make(chan int)
	loadChan(c)
	for v := range c {
		fmt.Println(v)
	}
}

func loadChan(c chan int) chan int {
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}
