package main

// lesson 4: Arrays and Slices

import "fmt"

func main() {
	// declaring array with a size
	grades := [3]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)

	// declaring array which is just enough to hold the elements specified
	grades1 := [...]int{95, 87, 88, 76}
	fmt.Printf("Grades: %v\n", grades1)

	// declaring an empty array
	var students [4]string
	fmt.Printf("Students: %v\n", students)

	// indexing into the element
	students[0] = "Lisa"
	students[2] = "Arnold"
	students[1] = "Ahmed"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student at index 3: %v\n", students[3])
	fmt.Printf("Number of students in the array: %v\n", len(students))

	// 2D arrays 
	//var identityMatrix [3][3]int = [3][3]int {[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	var identityMatrix [3][3]int = [3][3]int {{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Println(identityMatrix)

	// other way to assign values in an array
	var idMatrix [2][2]int
	idMatrix[0] = [2]int{1, 0}
	idMatrix[1] = [2]int{0, 1}

	// in Go, the array are value types i.e. when we assign one array to another it creates a literal copy
	// i.e. they are different memory altogether
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5 // 'a' remains unchanged after this operation
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

	// to pass arrays by reference, we'll pass them using pointer
	c := [...]int{1, 3}
	d := &c
	d[1] = 7
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)

	// slices are just expandable arrays i.e. ArrayList in Java
	a1 := []int{1, 2, 3, 4}
	fmt.Printf("Length of a1: %v\n", len(a1))
	fmt.Printf("Capacity of a1: %v\n", cap(a1))

	// slices are reference types
	b1 := a1
	b1[3] = 6
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("b1: %v\n", b1)

	// creating slices by slicing operations (slices can also be created from an array the same way)
	// slices created from an array are still pass by reference and point to the underlying array
	// hence they will reflect change that are being made on the underlying array
	a2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b2 := a2[:] // slice of all elements
	c2 := a2[3:] // slice from 4th element to end
	d2 := a2[:6] // slice of first 6 elements
	e2 := a2[3:6] // slice the 4th, 5th and 6th element
	a2[5] = 42 // changes all the slices
	fmt.Println(a2)
	fmt.Println(b2)
	fmt.Println(c2)
	fmt.Println(d2)
	fmt.Println(e2)

	// creating a slice by specifying the size of the slice and the type
	a3 := make([]int, 3)
	fmt.Printf("Length of a3: %v\n", len(a3))
	fmt.Printf("Capacity of a3: %v\n", cap(a3))

	// creating a slice by specifying the size (append those many zero element of that type) and the capacity
	a4 := make([]int, 4, 100)
	fmt.Printf("Length of a4: %v\n", len(a4))
	fmt.Printf("Capacity of a4: %v\n", cap(a4))
}