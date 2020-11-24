package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	// exerciseOne()
	// exerciseTwo()
	// exerciseThree()
	// exerciseFour()
	exerciseFive()
}

func exerciseOne() {
	fmt.Println("=====||EXERCISE #1||=====")
	a := 42
	b := "James Bond"
	c := true

	// Single print statement
	fmt.Printf("%v\n%v\n%v\n", a, b, c)
	//Multiple print statements
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func exerciseTwo() {
	fmt.Println("=====||EXERCISE #2||=====")
	//Print out default zero values for variable types
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
}

func exerciseThree() {
	fmt.Println("=====||EXERCISE #3||=====")
	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)
}

type newint int

var n newint

func exerciseFour() {
	fmt.Println("=====||EXERCISE #4||=====")
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	n = 42
	fmt.Println(n)
}

var m int

func exerciseFive() {
	fmt.Println("=====||EXERCISE #5||=====")
	exerciseFour()
	m = int(n)
	fmt.Println(m)
	fmt.Printf("%T\n", m)
}
