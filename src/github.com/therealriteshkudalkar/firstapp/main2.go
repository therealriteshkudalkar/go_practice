package main

// lesson 2: Type Conversion

import (
	"fmt"
	"strconv"
)

var (
	actorName string = "Elizabeth Sladen"
	companion string = "Sarah Jane Smith"
	doctorNumber int = 2
)

var (
	counter int = 0
	index int = 0
	i int = 56
)

func main2() {
	fmt.Println("Hello, world!")

	fmt.Println(42)

	var i int = 32
	fmt.Println(i)

	j := 42
	fmt.Println(j)

	k := 99.25
	fmt.Printf("%v, %T\n", k, k)
	
	var n float32 = 65.43
	m := float64(n)
	fmt.Printf("%v, %T\n", m, m)

	strn := strconv.FormatFloat(m, 'f', 6, 64)
	fmt.Printf("%v, %T\n", strn, strn)
}
