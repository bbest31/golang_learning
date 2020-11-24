package main

import "fmt"

type person struct {
	first            string
	last             string
	favoriteIceCream []string
}

var p person = person{
	first:            "Brandon",
	last:             "Best",
	favoriteIceCream: []string{"Salted Caramel", "Netflix N' Chill", "Lemon"},
}

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
}

func ex1() {
	fmt.Println("=====|| 1 ||=====")

	p1 := person{
		first:            "Brandon",
		last:             "Best",
		favoriteIceCream: []string{"Salted Caramel", "Netflix N' Chill", "Lemon"},
	}

	p2 := person{
		first:            "Emma",
		last:             "Wickenheiser",
		favoriteIceCream: []string{"Hazelnut", "Netflix N' Chill", "Strawberry"},
	}

	for k, v := range p1.favoriteIceCream {
		fmt.Println(k, v)
	}

	for k, v := range p2.favoriteIceCream {
		fmt.Println(k, v)
	}
}

func ex2() {
	fmt.Println("=====|| 2 ||=====")
	m := map[string][]string{}
	m[p.last] = p.favoriteIceCream

	fmt.Println(m)

	for k, v := range m {
		for _, iceCream := range v {
			fmt.Println(k, iceCream)
		}
	}
}
func ex3() {
	fmt.Println("=====|| 3 ||=====")
	type vehicle struct {
		doors int
		color string
	}

	type sedan struct {
		vehicle
		luxury bool
	}
	type truck struct {
		vehicle
		fourWheel bool
	}

	cybertruck := truck{
		vehicle: vehicle{
			doors: 4,
			color: "Stainless Steel",
		},
		fourWheel: true,
	}

	roadster := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "Black",
		},
		luxury: true,
	}

	fmt.Println(cybertruck)
	fmt.Println(roadster)
	fmt.Println(cybertruck.color)
	fmt.Println(roadster.luxury)
}
func ex4() {
	fmt.Println("=====|| 4 ||=====")

	h := struct {
		bedrooms  int
		bathrooms int
		sqft      int
	}{
		bedrooms:  2,
		bathrooms: 2,
		sqft:      800,
	}

	fmt.Println(h)

}
