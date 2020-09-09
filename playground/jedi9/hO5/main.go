package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//using atomic instead of a mutex
func main() {
	var counter int64
	counter = 0
	var wg sync.WaitGroup
	wg.Add(100)
	//var tex sync.Mutex
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)

			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
