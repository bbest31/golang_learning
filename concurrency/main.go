package main

import (
	"fmt"
	"runtime"
	"sync"
)

// func init() {} runs before func main()

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine())

	wg.Add(1) // Adds a Go routine to the Wait group
	go foo()  // keyword go launches the foo() in another Goroutine, however the main func doesn't wait for it to finished without some other parts.
	bar()

	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())
	wg.Wait() // Point at which the program will wait for any Goroutines inside the WaitGroup.
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done() // Call to tell the WaitGroup that this Goroutine is done.
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
