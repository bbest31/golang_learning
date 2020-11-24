package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go func() {
		fmt.Println("Method 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("Method 2")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Main function exiting...")
}
