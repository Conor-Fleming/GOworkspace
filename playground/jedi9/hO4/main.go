package main

import (
	"fmt"
	"sync"
)

//fixing a race condition by using a mutex to lock a variable while in use
func main() {
	counter := 0
	var wg sync.WaitGroup
	wg.Add(100)
	var tex sync.Mutex
	for i := 0; i < 100; i++ {
		go func() {
			tex.Lock()
			v := counter
			//runtime.Gosched()
			v++
			counter = v
			tex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
