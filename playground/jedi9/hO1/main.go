package main

import (
	"fmt"
	"sync"
)

var waitG sync.WaitGroup

func main() { // using wait groups to make sure go routines finish
	waitG.Add(3)
	go go1()
	go go2()
	go go3()
	bar()
	waitG.Wait()
}

func bar() {
	fmt.Println("basic routine")
}

func go1() {
	fmt.Println("this is routine 1")
	waitG.Done()
}

func go2() {
	fmt.Println("this is routine 2")
	waitG.Done()
}

func go3() {
	fmt.Println("this is routine 3")
	waitG.Done()
}
