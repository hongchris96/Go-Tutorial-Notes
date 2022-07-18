package main

import (
	"fmt"
)

func main() {

	// Defer
	fmt.Println("Defer:")
	// the function that defer takes in is called after the main function

	fmt.Println("one")
	defer fmt.Println("two")   // this line gets moved after the main function but before the main function returns
	defer fmt.Println("three") // defer order works using LIFO: will print four, three, two
	defer fmt.Println("four")
	fmt.Println("five")

	a := "a value before defer print"
	defer fmt.Println(a) // function in defer takes the value of a at the moment
	a = "a value after defer print"

}
