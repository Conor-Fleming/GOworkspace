package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:\t", runtime.NumCPU())
	fmt.Println("GRs:\t", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	var mu sync.Mutex //using mutex to prevent race condition
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GRs:\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println(runtime.NumCPU(), runtime.NumCPU())
	fmt.Println(counter)
}
