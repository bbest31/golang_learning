package main

import (
	"fmt"
	"math"
)

// When to use pointers
// 1. when you have a big piece of data you can pass around the address instead of the data itself.
// 2.
func main() {
	a := 31
	fmt.Println(a)
	fmt.Println(&a)        // & operator to view address
	fmt.Printf("%T\n", &a) // type is a pointer to an int = *int

	var p *int = &a
	fmt.Println(p)
	fmt.Println(*p) // de-references the address to access the value at that address

	*p = 32 // changes value at that address
	fmt.Println(a)

	x := 0
	foo(x)
	fmt.Println(x) // x not changed even when we tried to reassign in foo()
	bar(&x)
	fmt.Println(x) // x changed cause we passed a pointer

	c := circle{5}
	info(&c) // has to be a pointer passed b/c the implemented method in circle uses a pointer in the reciever.
}

func foo(y int) {
	fmt.Println(y)
	y = 1
	fmt.Println(y)
}

func bar(y *int) {
	fmt.Println(*y)
	*y = 1
	fmt.Println(*y)
}

type circle struct {
	r float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 { // this function has a pointer reciever. So it can only be called on a pointer to a circle.
	return math.Pi * c.r * c.r
}

func info(s shape) {
	fmt.Println(s.area())
}
