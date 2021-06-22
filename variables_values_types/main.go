package main // each application needs one file with

import (
	"fmt"
	"reflect"
)

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

	// Can only do type switching with interfaces
	// Use the reflect package instead
	//.TypeOf returns general types
	// .TypeOf(x).Kind() returns specific types like Int, Int8, Int16, etc.

	if reflect.TypeOf(x) == reflect.TypeOf(y) {
		fmt.Println("x is same type as y")
	}

}

func isInt(val interface{}) bool {
	if _, ok := val.(int); ok {
		return true
	}

	return false
}

func printType(x interface{}) {
	// Comparing types using a type switch
	switch x.(type) {
	case string:
		fmt.Println(x, "(string)")
	case float64:
		fmt.Println(x, "(float64)")
	case int:
		fmt.Println(x, "(interface)")
	default:
		fmt.Printf("(%T)\n", x)
	}
}
