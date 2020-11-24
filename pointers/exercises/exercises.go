package main

import "fmt"

func main() {
	fmt.Println("=====|| #1 ||=====")
	x := 31
	fmt.Println(&x)

	fmt.Println("=====|| #2 ||=====")

	p := person{"Brandon", "Best"}
	fmt.Println(p)

	changeMe(&p)
	fmt.Println(p)
}

type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	*p = person{"The", "Chrome"} // if you want to change the whole thing
	// p.last = "B." // Valid way to change fields if p is of pointer type
}
