package main

import "fmt"

// Structs are actually composite types
type spaceMarine struct {
	chapter string
	rank    int
	weapon  //anonymous field which is named by the type given
}

// We can also embed structs
type weapon struct {
	name  string
	melee bool
	Range bool
}

func main() {
	marine1 := spaceMarine{
		chapter: "Grey Knights",
		rank:    4,
		weapon: weapon{
			name:  "Force Halberd",
			melee: true,
			Range: true,
		},
	}
	marine2 := spaceMarine{
		chapter: "Blood Angels",
		weapon: weapon{
			name:  "Bolter",
			melee: false,
			Range: true,
		},
	}

	fmt.Println(marine1)
	fmt.Println(marine2)

	fmt.Println(marine1.chapter)
	fmt.Println(marine1.weapon)
	fmt.Println(marine1.melee) // can bypass the marine1.weapon.melee if the attribute is an anonmyous field, the melee, name, and Range fields are considered to be promoted in this case.

	fmt.Println(marine2.rank) // uninitialized values get zero values based on type

	// Anonymous Structs
	// structs where you may only need it in one spot so just create it inline
	calth := struct {
		name    string
		gravity float64
		orbit   float64
	}{
		name:    "Calth",
		gravity: 1.456,
		orbit:   0.6,
	}

	fmt.Println(calth)
}
