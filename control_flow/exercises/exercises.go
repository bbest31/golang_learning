package main

import "fmt"

func main() {
	// exc1()
	// exc2()
	// exc3()
	// exc4()
	// exc5()
	// exc6(3)
	// exc7(0)
	// exc8()
	// exc9()
	// exc10()

}

func exc1() {
	fmt.Println("=====|| #1 ||=====")
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}
}

func exc2() {
	fmt.Println("=====|| #2 ||=====")
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}

func exc3() {
	fmt.Println("=====|| #3 ||=====")
	y := 1994
	for y <= 2020 {
		fmt.Println(y)
		y++
	}
}

func exc4() {
	fmt.Println("=====|| #4 ||=====")
	y := 1994
	for {
		if y > 2020 {
			break
		}
		fmt.Println(y)
		y++
	}
}

func exc5() {
	fmt.Println("=====|| #5 ||=====")
	for i := 0; i <= 100; i++ {
		fmt.Printf("%v %% 4 = %v\n", i, i%4)
	}
}

func exc6(n int) {
	fmt.Println("=====|| #6 ||=====")
	if n%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}

func exc7(n int) {
	fmt.Println("=====|| #7 ||=====")
	if n == 0 {
		fmt.Println("zero")
	} else if n < 0 {
		fmt.Println("negative")
	} else {
		fmt.Println("positive")
	}
}

func exc8() {
	fmt.Println("=====|| #8 ||=====")
	switch {
	case 0 > 1:
		fmt.Println("should not print")
	case 0 < 1:
		fmt.Println("should print")
	}
}

func exc9() {
	fmt.Println("=====|| #9 ||=====")
	favSport := "Basketball"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("go to the pool!")
	case "surfing":
		fmt.Println("go to hawaii!")
	case "wingsuit flying":
		fmt.Println("what would you like me to say at your funeral?")
	case "Basketball":
		fmt.Println("Go to the courts!")
	}
}

func exc10() {
	fmt.Println("=====|| #10 ||=====")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
