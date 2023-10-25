package main

// lesson 8: Control Flow with defer, panic and recover

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// the defer keyword executes the function or block of code or statement passed into it
	// after the execution of current function in which it is used and before the current
	// function returns the control to the caller function
	fmt.Println("Start")
	defer fmt.Println("Middle")
	fmt.Println("End")

	demo()

	demo1()

	demo2()

	demo3()

	demo4()

	// panics can also affects the control flow of the program
	//a, b := 1, 0
	//ans := a / b
	//fmt.Println(ans)

	// we can also panic on our own
	fmt.Println("before panic")
	//panic("Something bad happened!")
	fmt.Println("after panic")

	//demo5()

	// control flow with panics and defer statement
	// all the defer statements are executed before any panic and then return happens
	fmt.Println("Start the job")
	defer fmt.Println("This was deferred before panic.")
	//panic("I'm panicking!!")
	fmt.Println("End the job")

	// a thing to note about the defer statement is that it takes in a function call and
	// not the function itself
	fmt.Println("Start something")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("Something bad happened after the defer.")
	fmt.Println("End something")

	// recover is useful when we have a deeper call stack than this
}

func demo() {
	// if there are multiple defer statements then the calls are execute in LIFO manner
	defer fmt.Println("Start Demo")
	defer fmt.Println("Middle Demo")
	defer fmt.Println("End Demo")
}

func demo1() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := io.ReadAll(res.Body)
	// after reading close the body
	// but what if we forget to close the body somewhere down the line?
	// so instead use defer and use Body.Close() method just after the res statement
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Robots: %v\n", string(robots))
}

func demo2() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}

	//defer allow you to write the code opening and closing a resource together
	// defer just after handling the error
	defer res.Body.Close()

	// point to note is that defer will not work when we are opening a resource in a loop
	// it will keep on opening resources and those will only be closed when the current function ends

	robots, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Robots: %v\n", string(robots))
}

func demo3() {
	// the defer keyword takes the input right where it is called
	a := "start demo3"
	defer fmt.Println(a)
	a = "end demo3"
}

type Cat struct {
	name string
	age  int
}

func alter_struct(cat *Cat) {
	cat.name = "Kaitlin"
	fmt.Println("Altered cat: ", *cat)
}

func demo4() {
	// a very stranger behaviour with defer statement
	a := Cat{
		name: "Manu",
		age:  55,
	}
	fmt.Println("Before alteration: ", a)
	defer fmt.Println("After second alteration: ", &a) // here the struct is passed by reference
	defer fmt.Println("After second alteration: ", a)  // here the struct is passed by value (so it takes stale value)
	defer alter_struct(&a)
	a.age = 70
	fmt.Println("After first alteration: ", a)
}

func demo5() {
	// create an http web application
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}
