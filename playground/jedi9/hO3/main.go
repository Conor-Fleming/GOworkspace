package main

import (
	"fmt"
	"runtime"
	"sync"
)

//creating a race condition
func main() {
	counter := 0
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			//fmt.Println(counter)
			wg.Done()
		}()
		//fmt.Println(counter)
	}
	fmt.Println(counter)
	wg.Wait()
}
