package main

// lesson 6: Control Flow

import (
	"fmt"
	"math"
)

func main6() {
	// the curly braces are necessary even if there is only one statement inside the body
	if true {
		fmt.Println("The test is true.")
	}

	// the condition specified in the if statement are always suppose to be of the type boolean
	// if 0 {
	// 	fmt.Println("Gives a compiler error!")
	// }

	statePopulation := map[string]int{
		"California":   39_250_017,
		"Texas":        27_862_596,
		"Florida":      20_612_439,
		"New York":     19_745_289,
		"Pennsylvania": 12_802_503,
		"Illinois":     12_801_539,
		"Ohio":         11_614_373,
	}

	// if with an initializer syntax
	// here the variables initialized are only scoped to the if block and cannot be used outside
	if pop, ok := statePopulation["Florida"]; ok {
		fmt.Println(pop)
	}

	// pop is not available here
	//fmt.Printlnl(pop)

	// various comparison operators that result in the bool
	number := 50
	guess := 101
	if guess < number {
		fmt.Println("Too low")
	}
	if guess > number {
		fmt.Println("Too high")
	}
	if guess == number {
		fmt.Println("You got it!")
	}
	fmt.Println(number <= guess, number >= guess, number != guess)

	// we can also check for string equality
	s1 := "hello"
	s2 := "hey"
	s3 := "hey"
	s4 := "hella"
	if s1 == s2 {
		fmt.Println("The strings s1 and s2 are equal")
	}
	if s2 == s3 {
		fmt.Println("The strings s2 and s3 are equal.")
	}

	// checks if s1 comes lexicographically later than s1
	if s1 > s4 {
		fmt.Println("Yes, hello comes after hella.")
	}

	// logical operator
	if guess < 1 || guess > 100 { // logical OR operator
		fmt.Println("The guess must be between 1 and 100!")
	}
	if guess >= 1 && guess <= 100 { // logical AND operator
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it!!")
		}
	}
	fmt.Println(!true) // logical NOT operator

	// shortcircuiting also takes place in Go
	// it does not execute the returnsTrue function if it finds that atlease one operand in
	// the logical OR is true
	if false || returnsTrue() {
		fmt.Println("The if statement's condition is true: body")
	}
	// it does not execute the returnsTrue function if it finds that atleast one operand in
	// the logical AND is false
	if false && returnsTrue() {
		fmt.Println("The if statement's condition is true: body")
	}

	// if else-if and else
	if guess >= 1 && guess <= 100 {
		if guess < number {
			fmt.Println("Too low")
		} else if guess > number {
			fmt.Println("Too high")
		} else {
			fmt.Println("You found it!")
		}
	} else {
		fmt.Println("The guess must be between 1 and 100 inclusive.")
	}

	// floating point numbers are approximations hence we
	// should not use them in comparison operations
	myNum := 0.123
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("These are the same!")
	} else {
		fmt.Println("These are different!")
	}

	// to check if the numbers are same, we can use how close the numbers are
	// by dividing the numbers
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.0001 {
		fmt.Println("These are the same!")
	} else {
		fmt.Println("These are different!")
	}

	// switch statement
	num := 3
	switch num { // here num variable is called the switch's tag
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("not one or two or three")
	}

	// fallthrough is not a default behaviour in Go's switch statement
	// we cannot have overlapping cases in switch
	switch num {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	}

	// in switch we are not limited to boolean value but other values types
	// and just like if else we can use the initilizer syntax by scoping the tag variable
	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	}

	// we can also have a tagless switch in which
	// in tagless switch, we can have overlapping cases
	// but the first case which matches the tag variable is executed
	// the break is implicit in switch
	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
	case i <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

	// we can also explicitly fallthrough
	// all the statements are executed after fallthrough regardless of them passing the condition or not
	i = 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough
	case i <= 20:
		fmt.Println("less than or equal to twenty")
		fallthrough
	case i >= 30:
		fmt.Println("greater than thirty") // this is executed even though i is less than 30
	default:
		fmt.Println("greater than twenty")
	}

	// we can also have a type switch
	// we can use break keyword to break our of a switch case early
	// like if statement, we can also use initilizer syntax with switch
	var j interface{} = 34
	switch jType := j.(type) {
	case int:
		fmt.Println("i is an int")
		if jType < 10 {
			break
		}
		fmt.Println("This will print too")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is a string")
	case [3]int:
		fmt.Println("i is a [3]int")
	default:
		fmt.Println("i is another type")
	}
}

func returnsTrue() bool {
	fmt.Println("The return true function has executed.")
	return true
}
