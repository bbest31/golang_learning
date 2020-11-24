package main

import (
	"fmt"
	"math"
)

func main() {

	s := square{5.52}
	c := circle{3.74}

	info(s)
	info(c)
}

type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func ex1() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 1
}

func bar() (int, string) {
	return 1, "one"
}
