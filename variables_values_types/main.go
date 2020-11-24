package main // each application needs one file with

import "fmt"

// Keyword var is used to declare variables outside function bodies and be used accross them.
// Take care to limit these and use function scope variables as much as possible.
var y = 100

// declared variable z of type int with the ZERO VALUE of the int type.
//Zero Values int = 0, boolean = false, 0.0 = floats, "" = strings, nil = pointers, functions, interfaces, slices, channels, and maps
var z int

// Custom types - this declares a new type with an underlying base type.
type bitcoin float64

var dogecoin bitcoin

func main() {
	x := 42 // Short declaration operator := declares a variable and sets the value.
	fmt.Println("Hello, world!", x)
	x = 99 // once a variable is declared we can just use = to reassign it.
	fmt.Println(y)
	// print the type of z
	fmt.Printf("%T\n", z)
	fmt.Printf("%T\n", dogecoin)
	fmt.Printf("Converting the type of a variable: %T\n", int(dogecoin))
}
