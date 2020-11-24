package main

import "fmt"

func main() {
	// Arrays
	var arr [5]int
	fmt.Println(arr)
	// access and change elements
	arr[3] = 31
	fmt.Println(arr)
	fmt.Println(len(arr))

	// Slices - should be used over Arrays
	// They are composite literals where you can group together values of the same type.
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)

	// loop over slice
	for index, value := range x {
		fmt.Println(index, value)
	}

	// slicing a Slice
	fmt.Println(x[0:3])

	//appending
	fmt.Println(append(x, 6, 7, 8, 9))

	y := []int{12, 13, 14, 15}

	x = append(x, y...)
	fmt.Println(x)
	// delete from a slice
	x = append(x[:2], x[3:]...)
	fmt.Println(x)

	// Can use make to create a Slice to give the underlying array a specific size to avoid an early array expansion procedure.
	x = make([]int, 10, 12) // length of the slice is 10, and the capacity is 100
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, -1, -3, 0) // appending beyond the cap will prompt a doubling in cap space by copying original underlying array to new array of double capacity.
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// Iterate over a Slice
	for i, v := range x {
		fmt.Println(i, v)
	}

	// Multi-dimensional slicing

	gk := []string{"Grey Knights", "Titan", "Ordo Malleus"}
	fmt.Println(gk)
	da := []string{"Dark Angels", "Caliban", "Adeptus Astartes"}
	fmt.Println(da)

	z := [][]string{gk, da}
	fmt.Println(z) // indexing in mulit-dimensional array z[i][j]...[n]

	// Maps
	//map[key type]value type{}
	m := map[string]int{
		"Dark Angels":       1,
		"Iron Warriors":     9,
		"Sons of Horus":     16,
		"World Eaters":      12,
		"Iron Hands":        10,
		"Emperors Children": 3,
		"Raven Guard":       19,
	}
	fmt.Println(m)
	fmt.Println(m["Iron Hands"])
	v, ok := m["Salamanders"]
	fmt.Println(v)  // 0
	fmt.Println(ok) // tells you if the value exists in the map

	// Add to map
	m["White Scars"] = 5

	// iterate over a map
	for k, v := range m {
		fmt.Println(k, v)
	}

	// Delete from Map

	delete(m, "Sons of Horus")
	fmt.Println(m)
	// can use the ok idiom to ensure there is an item to delete with the corresponding key.
	if v, ok := m["Word Bearers"]; ok {
		fmt.Println(v)
		delete(m, "Word Bearers")
	} else {
		fmt.Println("That k:v doesn't exist")
	}

}
