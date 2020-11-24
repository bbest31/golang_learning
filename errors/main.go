package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	// _, err := os.Open("no-file.txt")
	// if err != nil {
	// fmt.Println("err happended", err)
	// log.Println("err happended", err)
	// log.Fatalln(err)
	// panic(err)
	// }

	// Defer statement
	// The function arguments are evaluated as they were when this statement is written. Therefore i will be 0 not 1.
	// Deferred function calls are executed in LIFO (stack)
	i := 0
	defer fmt.Println(i)
	i++
	// Deferred functions can also alter the return values if they are incorrect
	defer func() { i++ }()

	// Recover
	// built-in function to regain control of a panicking goroutine.
	// Only useful inside deferred functions
	foo()

	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func foo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in foo", r)
		}
	}()
	fmt.Println("Calling bar")
	bar(0)
	fmt.Println("Returned normall from bar().")
}

func bar(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in bar", i)
	bar(i + 1)
}

// Fatalln is equivalent to Println followed by a call to os.Exit(1),  and deferred functions are not run
// Panic is similar to Fatalln but deferred functions are run.

//NOTE: Suggested to use log package for flexibility on writing errors to different outputs.

// Ex. using errors.New

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgate math: square root of negative number")
		// return 0, fmt.Errorf("math err: square root of negative number: %v", f)
	}
	return 26, nil
}

// Custom Error

type myMathError struct {
	lat  string
	long string
	err  error
}

func (n myMathError) Error() string {
	return fmt.Sprintf("A math error ocurred: %v %v %v", n.lat, n.long, n.err)
}
