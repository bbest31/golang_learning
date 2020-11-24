package main

import "fmt"

func main() {

	bar("Brandon")
	s := foo("Brandon")
	fmt.Println(s)
	quotient, remainder := divide(12, 5)
	fmt.Printf("Quotient: %v\t Remainder: %v\n", quotient, remainder)

	total := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(total)

	evenTotal := even(sum, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(evenTotal)

	deferral()

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	sa1.shoot()

	// Anonymous function
	func() {
		fmt.Println("Anonymous function!")
	}() // parens calls it

	// Func expression
	f := func() {
		fmt.Println("Expression function!")
	}

	f()

	b := returnFx()

	fmt.Println(b())

	//Closure
	// can use a code block to limit scope of a variable

	clos := func() func() int {
		var x int // this is scoped to the contained function
		return func() int {
			x++
			return x
		}
	}()

	fmt.Println(clos())
	fmt.Println(clos())
	fmt.Println(clos())

	{
		var y int
		fmt.Println(y) // Works!
	}

	// fmt.Println(y) y is not defined outside the code block!

}

// Syntax
// func (r receiver) identifier(params) (return(s)) {...}

func bar(s string) {
	fmt.Println("Hello ", s)
}

func bar2(s string) {
	fmt.Println("Hello from ", s)
}

func foo(s string) string {
	return fmt.Sprint("Hello from foo, ", s)
}

// Multi-return
func divide(a int, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

// Variadic Parameters
// accept unlimited number of params or none at all
func sum(x ...int) int {
	total := 0
	fmt.Println(x)
	for _, v := range x {
		total += v
	}

	return total
}

// Defer statement
// A deferred function only runs once the containing body reaches it's end

func deferral() {
	// bar will be run first instead of foo
	defer bar("Brandon") // will run once deferral() reaches it's end
	bar2("Brandon")
}

// Methods
// attache a method to a type

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// Define the type that is allowed to access this method.
func (s secretAgent) shoot() {
	fmt.Printf("%v fired a shot!\n", s.first)
}

// Interfaces and Polymorphism

// Anything that has the method shoot() is now also of type human
// If there was more than one method then the subclass has to implement every interface method.
type human interface {
	shoot()
}

func (p person) shoot() {}

// Now both person and secretAgent types can be passed to this function because they both inherit the interfce type human
func sayName(h human) {
	h.shoot()
}

// Assertions
// asserting a value is of a specific type allows you to access type specfic fields
func switchOnType(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I am a person", h.(person).first) // this uses an assertion to make sure h is of type person
	case secretAgent:
		fmt.Println("I am a secret agemt", h.(secretAgent).ltk)
	}
}

// Returning a function
// This is a function returnFx which returns a func which returns a string
func returnFx() func() string {
	return func() string {
		return "This is a string!"
	}
}

// Callbacks
func even(f func(xi ...int) int, s ...int) int {
	var evens []int
	for _, v := range s {
		if v%2 == 0 {
			evens = append(evens, v)
		}
	}

	return f(evens...)
}
