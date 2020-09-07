package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:\t", runtime.NumCPU())
	fmt.Println("GRs:\t", runtime.NumGoroutine())

	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			//runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("GRs:\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println(runtime.NumCPU(), runtime.NumCPU())
	fmt.Println(counter)
}
