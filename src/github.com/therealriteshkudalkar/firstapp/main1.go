package main

// lesson 1: Variables and Types

import (
	"fmt"
)

func main1() {
	var b bool = true
	fmt.Printf("%v, %T\n", b, b)

	test := 1 == 1
	fmt.Printf("%v, %T\n", test, test)
	test = 1 > 2
	fmt.Printf("%v, %T\n", test, test)
	test = 2 < 5
	fmt.Printf("%v, %T\n", test, test)
	test = 2 != 5
	fmt.Printf("%v, %T\n", test, test)

	var k int8 = 127
	fmt.Printf("%v, %T\n", k, k)

	var l uint64 = 25_000_000
	fmt.Printf("%v, %T\n", l, l)

	var m uint32 = 0b01010101
	fmt.Printf("%v, %T\n", m, m)
	
	var n uint32 = 0xffff_1234
	fmt.Printf("%v, %T\n", n, n)

	fmt.Println(m + n) // add
	fmt.Println(m - n) // sub
	fmt.Println(m * n) // mul
	fmt.Println(m / n) // div
	fmt.Println(m % n) // mod
	fmt.Println(m ^ 2) // xor
	fmt.Println(m &^ 2) // nand
	fmt.Println(m & 2) // and
	fmt.Println(m | 2) // or
	fmt.Println(m << 2) // left shift
	fmt.Println(m >> 2) // right shift

	p := 3.14
	p = 13.7e72
	p = 2.1E14
	fmt.Printf("%v, %T\n", p, p)

	var q complex128 = 1 + 2i
	fmt.Printf("%v, %T\n", q, q)
	fmt.Printf("%v, %v\n", imag(q), real(q))
	s := complex(5, 128)
	fmt.Printf("%v, %T\n", s, s)

	t := "this is a string"
	// t[3] = "u" // not possible because strings are immutable
	fmt.Printf("%v, %T\n", t, t)
	fmt.Printf("%v, %T\n", t[2], t[2])

	bs := []byte(t)
	fmt.Printf("%v, %T\n", bs, bs)

	// runes
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)
}