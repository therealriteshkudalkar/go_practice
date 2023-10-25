package main

// lesson 7: looping

import (
	"fmt"
)

func main7() {
	// there is only one loop statement in go i.e. for loop
	// scope of i is limited to the body of the for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// we can also manipulate the counter variable inside the loop body
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			i /= 2
		} else {
			i = 2*i + 1
		}
	}

	// in Go, i++ is a statement and not an expression
	k, l := 0, 1
	//k = l++ // invalid
	fmt.Println(k, l)

	// we can initilize two loop variables
	for i, j := 0, 0; i < 5 || j < 7; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	// all the three statements specified in the loop are optional
	// here i is scoped to the main function
	i := 2
	for ; i < 5; i++ {
		fmt.Println(i)
	}

	// this is just like a while loop
	i = 2
	for i < 5 {
		fmt.Println(i)
		i++
	}
	fmt.Println(i)

	// we can have an infinite loop by not specifiying anything
	// we can also use a break keyword to come out of the infinite for loop
	i = 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}

	// we also have continue statement
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// example of nested for loop
label:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				// break // breaks only the inner loop i.e. closest loop
				break label // breaks out of outer for loop
			}
		}
	}

	// working will collections
	s := []int{7, 8, 9}
	fmt.Println(s)

	// this works with slices, arrays
	for index, value := range s {
		fmt.Println(index, " ", value)
	}

	// it also works with maps
	statePopulation := map[string]int{
		"California":   39_250_017,
		"Texas":        27_862_596,
		"Florida":      20_612_439,
		"New York":     19_745_289,
		"Pennsylvania": 12_802_503,
		"Illinois":     12_801_539,
		"Ohio":         11_614_373,
	}
	for key, value := range statePopulation {
		fmt.Printf("Key: %v, value: %v\n", key, value)
	}

	// if don't want both key and value we can use underscore to skip
	for _, value := range statePopulation {
		fmt.Println(value)
	}

	// the for range loop also works with strings
	str := "Hello Go!"
	for index, value := range str {
		fmt.Printf("Index: %v, Character: %v\n", index, string(value))
	}
}
