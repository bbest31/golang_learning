package main

import "fmt"

func main() {
	ex1(31)
	ex2()
	ex3()
	ex4()
	ex5()
	ex6()
}

func ex1(x int) {
	fmt.Println("=====|| #1 ||=====")
	fmt.Printf("%v\t%b\t%X\n", x, x, x)
}

func ex2() {
	fmt.Println("=====|| #2 ||=====")
	a1 := 10 == 10 // true
	a2 := 90 <= 10 // false
	a3 := 0 >= -80 // true

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
}

const (
	c1     = 9.81
	c2 int = 80
)

func ex3() {
	fmt.Println("=====|| #3 ||=====")
	fmt.Println(c1)
	fmt.Println(c2)
}

func ex4() {
	fmt.Println("=====|| #4 ||=====")
	n := 100
	fmt.Printf("%v\t%b\t%X\n", n, n, n)
	nShifted := n << 1
	fmt.Printf("%v\t%b\t%X\n", nShifted, nShifted, nShifted)
}

func ex5() {
	fmt.Println("=====|| #5 ||=====")
	stringLiteral := `This is a string literal
	which will keep all the returns and 	tabs.`
	fmt.Print(stringLiteral)
}

func ex6() {
	fmt.Println("=====|| #6 ||=====")
	const (
		_  = iota
		y1 = 2020 + iota
		y2 = 2020 + iota
		y3 = 2020 + iota
		y4 = 2020 + iota
	)

	fmt.Println(y1)
	fmt.Println(y2)
	fmt.Println(y3)
	fmt.Println(y4)
}
