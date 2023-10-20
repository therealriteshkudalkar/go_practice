package main

// lesson 5: Maps and Structs

import (
	"fmt"
)

func main() {
	statePopulation := map[string]int{
		"California":   39_250_017,
		"Texas":        27_862_596,
		"Florida":      20_612_439,
		"New York":     19_745_289,
		"Pennsylvania": 12_802_503,
		"Illinois":     12_801_539,
		"Ohio":         11_614_373,
	}
	fmt.Printf("Population: %v\n", statePopulation)

	// for a certian type to be a key in a map, it we should be able to test it for equality
	// we cannot use slice, maps and functions as keys for a map

	// m := map[[]int]string{} // invalid

	// an array however can be used as a key in a map
	m := map[[3]int]string{}
	fmt.Printf("Empty Map: %v\n", m)

	// we can also create maps using the make function
	m1 := make(map[string]int)

	// indexing in the map
	m1["hello"] = 10
	m1["the"] = 250
	fmt.Printf("Count Map: %v\n", m1)
	fmt.Printf("hello count: %v\n", m1["the"])
	fmt.Printf("man count: %v\n", m1["man"])

	// deleting entries from the map
	delete(m1, "the")
	fmt.Printf("Count Map: %v\n", m1)
	fmt.Printf("man count: %v\n", m1["the"])

	// check for existence of a key in the map
	_, ok := statePopulation["Nevada"]
	fmt.Printf("Is Nevada present? %v\n", ok)

	// finding number of keys in the map
	fmt.Printf("Number of states: %v\n", len(statePopulation))

	//like slices, maps are also passed by reference
	sp := statePopulation
	delete(sp, "Ohio")
	fmt.Printf("sp: %v\n", sp)
	fmt.Printf("statePopulations: %v\n", statePopulation) // deletion is reflected back

	type Doctor struct {
		number     int
		actorName  string
		companions []string
	}

	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Je Cren",
			"Mina Shane",
		},
	}
	fmt.Printf("Doctor: %v\n", aDoctor)

	// dot syntax for accessing the members
	fmt.Printf("Doctor's companion: %v", aDoctor.companions[1])

	//
}
