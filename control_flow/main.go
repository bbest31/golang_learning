package main

import "fmt"

func main() {

	// For loops
	// for init; condition; post {}
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}

	// Can use a boolean expression as the loop condition
	a := 2
	b := 20
	for a <= b {
		fmt.Printf("a = %v\tb = %v\n", a, b)
		a *= 2
	}

	// IF statements
	if a%2 == 0 {
		fmt.Println("a is even")
	} else {
		fmt.Println("a is odd")
	}

	// Boolean operators
	// && - AND
	// || - OR
	// ! - NOT

	// preceed condition with statement
	// The declared variable is available in all branches
	if num := a * 100; num < 0 {
		fmt.Println("a is negative")
	} else if num < 100 {
		fmt.Println("a < 100")
	} else {
		fmt.Println("a > 100")
	}

	//SWITCH

	switch a {
	case 2:
		fmt.Println("2")
	case 20:
		fmt.Println("20")
	case 200:
		fmt.Println("200")
	case 2000:
		fmt.Println("2000")
	default:
		fmt.Println(a)
	}

	// capture a switch in a function
	switchCase := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Printf("unknown %T\n", t)
		}
	}

	switchCase(4)
	switchCase(true)
	switchCase("string")

}
