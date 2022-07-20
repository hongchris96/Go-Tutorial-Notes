package main

import (
	"fmt"
	"log"
)

func main() {

	// Defer
	fmt.Println("Defer:")
	// the function call that defer takes in is called after the main function

	fmt.Println("one")
	defer fmt.Println("two")   // this line gets moved after the main function but before the main function returns
	defer fmt.Println("three") // defer order works using LIFO: will print four, three, two
	defer fmt.Println("four")
	fmt.Println("five")

	a := "a value before defer print"
	defer fmt.Println(a) // function in defer takes the value of a at the moment
	a = "a value after defer print"

	// --------------------------------------------------------------------------

	// Panic
	fmt.Println("Panic:")
	// works like exception in other languages when to handle errors
	// panic message runs after the defer statement but before the main function returns
	// when panic happens it won't run any line after, but it will run defer functions

	// panic("custom panic error message")
	// fmt.Println("this won't run")
	panicker()
	fmt.Println("after panicker")

}

func panicker() {
	fmt.Println("Start of Panic")
	defer func() {
		// recover function: recover from panic, only useful in a defer function
		// returns error that's causing the app to panic
		// returns nil if there's no panic
		if err := recover(); err != nil {
			log.Println("Error:", err)
			panic(err) // re throw the panic if you can't handle the error
		}
	}()
	panic("panic line inside panicker")
	// this won't run
	// this function that recover still stops execution
	// but function higher up the callstack still works
	fmt.Println("End of Panic")
}
