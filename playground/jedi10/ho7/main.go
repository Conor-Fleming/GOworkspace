package main

import "fmt"

func main() {
	c := make(chan int)
	for y := 0; y < 10; y++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}

	for v := 0; v < 100; v++ {
		fmt.Println(<-c)
	}

}
