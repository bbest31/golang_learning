package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mtx sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mtx.Lock() // Use a mutex to lock access to the code in this Goroutine
			v := counter
			runtime.Gosched() // allows other goroutines to run by yielding the CPU.
			v++
			counter = v
			mtx.Unlock() // Unlock access to code so other Goroutines can use it again.
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("count:", counter)

	// Example using low level sync/atomic package. This is for very low-level meticulous usage of memory to avoid race conditions and optimization.

	var counter2 int64

	var wg2 sync.WaitGroup
	wg2.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter2, 1)                           // increment a Int64 atomic type variable
			fmt.Println("Counter 2\t", atomic.LoadInt64(&counter2)) // loads the Int64 atomic variable
			runtime.Gosched()
			wg2.Done()
		}()
	}

	wg2.Wait()
	fmt.Println("count:", counter)
}
