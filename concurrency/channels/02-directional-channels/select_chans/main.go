package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(eve, odd, quit)
	recieve(eve, odd, quit)
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	//close(e)
	//close(o)
	q <- 0
}

func recieve(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("it is even:", v)
		case v := <-o:
			fmt.Println("it is odd:", v)
		case v := <-q:
			fmt.Println("time to quit:", v)
			return
		}
	}
}
