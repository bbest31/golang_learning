package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type person struct {
	First string `json:"First"` // tags are used to define corresponding json key when marshalling. If left out is uses the field name from the struct
	Last  string `json:"Last"`  // if using `json:"Last,omitempty"` then that field will be omitted if empty in the struct
	Age   int    `json:"Age"`   // if using `json:"-"` always omit this field from the json object
}

type byAge []person

// In order to have a custom sort you must implement 3 methods

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {

	p1 := person{
		"Brandon",
		"Best",
		26,
	}

	p2 := person{
		"Emma",
		"Best",
		22,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	// Marshal structs or slices of structs into JSON
	byteSlice, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteSlice))

	//Unmarshal JSON data
	s := `[{"First":"Brandon","Last":"Best","Age":26},{"First":"Emma","Last":"Best","Age":22}]`
	bs := []byte(s)
	fmt.Printf("%T\n%v\n", bs, bs)

	var ppl []person

	err = json.Unmarshal(bs, &ppl)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ppl)

	// NOTE:
	/*
		You can also make a type that implements the Writer interface which can be used by the Encode and Decode methods instead of Marshal and Unmarshal.
		Encode and Decode instantly go from Go to JSON or vice-versa and if you pass it a type that impements Writer when you can have that type
		automatically respond to http requests or send a request without saving the conversion to a variable first.
	*/

	// Sorting
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	sort.Ints(xi)
	sort.Strings(xs)
	fmt.Println(xi)
	fmt.Println(xs)

	// Sorting structs
	sort.Sort(byAge(people))
	fmt.Println(people)
}
