package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {

	ex1()
	ex2()
	ex3()

}

func ex1() {
	fmt.Println("=====|| #1 ||=====")
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
}

func ex2() {
	fmt.Println("=====|| #2 ||=====")
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}
	bs, err := toJSON(p1)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(bs))
}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("Failed to Marshal Go struct to JSON object: %v", err)
	}
	return bs, nil
}

func ex3() {
	fmt.Println("=====|| #3 ||=====")

	err := customErr{"custom error message"}
	oops(err)
}

type customErr struct {
	message string
}

func (e customErr) Error() string {
	return fmt.Sprintf("Here is the error: %v", e.message)
}

func oops(e customErr) {
	fmt.Println("oops ran -", e)
}
