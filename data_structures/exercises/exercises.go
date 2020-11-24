package main

import "fmt"

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
	ex6()
	ex7()
	ex8()
	ex9()
	ex10()

}

func ex1() {
	fmt.Println("=====|| 1 ||=====")
	var arr [5]int
	arr[0] = -1
	arr[1] = -2
	arr[2] = -3
	arr[3] = -4
	arr[4] = -5
	fmt.Println(arr)
}

func ex2() {
	fmt.Println("=====|| 2 ||=====")
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for i, v := range s {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", s)
}

func ex3() {
	fmt.Println("=====|| 3 ||=====")
	s := []int{40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52}
	s1 := s[2:7]
	s2 := s[7:12]
	s3 := s[4:9]
	s4 := s[3:8]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}

func ex4() {
	fmt.Println("=====|| 4 ||=====")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}

func ex5() {
	fmt.Println("=====|| 5 ||=====")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[:3], x[6:]...)
	fmt.Println(y)
}

func ex6() {
	fmt.Println("=====|| 6 ||=====")
	x := []string{
		` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,
	}
	fmt.Println(len(x))
	fmt.Println(cap(x))
	for i := 0; i < len(x); i++ {
		fmt.Printf("index = %v , value = %v\n", i, x[i])
	}
}

func ex7() {
	fmt.Println("=====|| 7 ||=====")
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	x := [][]string{jb, mp}

	for _, v := range x {
		for _, d := range v {
			fmt.Println(d)
		}
	}
}

func ex8() {
	fmt.Println("=====|| 8 ||=====")
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	for _, v := range m {
		for i, data := range v {
			fmt.Println(i, data)
		}
	}
}

func ex9() {
	fmt.Println("=====|| 9 ||=====")
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	m["best_brandon"] = []string{"Technology", "Fitness", "Pizza"}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func ex10() {
	fmt.Println("=====|| 10 ||=====")
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	delete(m, "bond_james")
	for k, v := range m {
		fmt.Println(k, v)
	}
}
