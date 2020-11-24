package main

import (
	"fmt"
	"runtime"
)

var x uint
var v int
var y float64
var z string
var w bool

const a = 42            // untyped constant has some flexibility on re-assignment
const b float64 = 42.78 // type constant
const c = "James Bond"

/*
Iota is a special identifier that increments everytime it is used in a statement.
We can use this to create enums easily
*/
const (
	j = iota // 0
	k        // 1
	l        // 2
)

const (
	reset = iota // 0
	r            // 1
	t            // 2
)

// Create an enum type called Size

type Size uint8

const (
	small Size = iota
	medium
	large
	extraLarge
)

func main() {
	// Runtime object allows insights into Go workspace settings and hardware insights
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.Compiler)

	// Strings are actually a Slice (Go name for Array) of bytes
	s := "Hello Go!"
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)

	fmt.Println("small = ", small)
	fmt.Println("medium = ", medium)
	fmt.Println("large = ", large)
	fmt.Println("extraLarge = ", extraLarge)

	// Bit shifting
	// x << y - multiplies x by 2 y times
	// x >> y - divides x by 2 y times
	// ex. 8 << 3 = 8 * 2 * 2 * 2 = 64
	bit := 2
	fmt.Printf("%d\t\t%b\n", bit, bit)

	bit2 := bit << 1
	fmt.Printf("%d\t\t%b\n", bit2, bit2)

	bit3 := bit2 << 1
	fmt.Printf("%d\t\t%b\n", bit3, bit3)

	bit4 := bit3 << 3
	fmt.Printf("%d\t\t%b\n", bit4, bit4)

}
