package main

import (
	"fmt"
)

// exported
const HeyConst int = 5

// not exported
const helloConst int16 = 10

func main3() {
	// Naming convention
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)

	// not valid, gives compile time error, as expressions that evaluate at runtime cannot be assigned to constanst
	//const newConst float64 = math.Sin(1.57)

	// valid as it evaluates to the same value and is not an output of the function call
	const newConst float64 = 65.5 * 25
	fmt.Printf("%v, %T\n", newConst, newConst)

	const a int = 14
	// invalid as constant needs to be assigned some value
	//const b string
	const b string = "foo"
	const c float32 = 3.14
	const d bool = true

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)

	// invalid, as we cannot declare collection types as constant
	// const e [3]string = [3]string{"hello", "mellow", "yellow"}

	// shadowing is possible with constant, also notice that we can change the type as well
	const helloConst int = 25
	fmt.Printf("%v, %T\n", helloConst, helloConst)

	// we shadow and make them as var also
	var HeyConst int = 35
	fmt.Printf("%v, %T\n", HeyConst, HeyConst)

	// constants can be added to variables
	const f int = 43
	var g int = 47
	fmt.Printf("%v, %T\n", f+g, f+g)

	// for untyped constants the type is inferred based on the expression the constant is used in
	// untyped constants behave like #define in C i.e. it replaces every instance with the value
	// untyped constants are flexible and can work with other types as well
	const h = 42 // inferred type depend on
	var i int16 = 27
	fmt.Printf("%v, %T\n", h+i, h+i)
	var j int64 = 85
	fmt.Printf("%v, %T\n", h+j, h+j)

	// enumerated constants
	const (
		x = iota
		y = iota
		z = iota
	)
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)
	fmt.Printf("%v, %T\n", z, z)

	// if we don't specify iota when it infers the pattern of assignment
	// the iota value is scoped to the constant block
	// so the counter starts from 0 on every block
	const (
		x1 = iota
		y1
		z1
	)
	fmt.Printf("%v, %T\n", x1, x1)
	fmt.Printf("%v, %T\n", y1, y1)
	fmt.Printf("%v, %T\n", z1, z1)

	const (
		errorSpecialist = iota
		catSpecialist
		dogSpecialist
		snakeSpecialist
	)
	var specialistType int = catSpecialist
	fmt.Printf("%v\n", specialistType == catSpecialist)

	// _ is a special type of variable in that it is a write only variable
	const (
		_ = iota + 5
		pathologist
		doctor
		nurse
	)
	var personOccupation int
	fmt.Printf("%v\n", personOccupation == pathologist)
	fmt.Printf("%v\n", pathologist)

	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)

	fileSize := 4_000_000_000.
	fmt.Printf("%.2fGB\n", fileSize/GB)

	const (
		isAdmin = 1 << iota
		isHeadquarters
		canSeeFinancials

		canSeeAfrica
		canSeeAsia
		canSeeEurope
		canSeeNorthAmerica
		canSeeSouthAmerica
	)
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHeadquarters&roles == isHeadquarters)
}
