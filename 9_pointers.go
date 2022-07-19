package main

import (
	"fmt"
)

func main() {
	// time start: 4:04:03 https://youtu.be/YS4e4q9oBaU

	// Pointers
	fmt.Println("Pointers:")

	// value types during copying doesn't point to the same memory
	a := 42
	b := a
	fmt.Println(a, b)
	a = 27
	fmt.Println(a, b)

	// making b a pointer to integer using long form syntax
	var newA int = 42
	// *: here is declaring newB a pointer by using * before type
	// &: address operator placed before newA
	var newB *int = &newA
	// newB is now holding the memory location of newA
	fmt.Println(newA, newB)
	fmt.Println(&newA, newB)
	// Dereferencing operator: putting * in front of a pointer
	fmt.Println(newA, *newB)
	// now both newA and newB are tied together
	newA = 27
	fmt.Println(newA, *newB)
	*newB = 14
	fmt.Println(newA, *newB)

	// Pointer arithmetic: (NOT available in Go, unless using the unsafe package)
	// doing math with memory addresses
	x := [3]int{1, 2, 3}
	y := &x[0]
	z := &x[1]
	fmt.Println("%v %p %p\n", x, y, z) // %p prints the value of pointer of y and z, which are addresses here
	// each element in an array are 4 bytes apart, so their memory address number has a difference of 4

	var ms *myStruct        // pointer to a myStruct
	fmt.Println(ms)         // shows nil, default pointer initialization value
	ms = &myStruct{foo: 42} // address of the object
	fmt.Println(ms)         // shows &{42}

	ms = new(myStruct) // new initialization syntax, gives foo a value of 0
	fmt.Println(ms)

	var ns *myStruct
	ns = new(myStruct)
	(*ns).foo = 42 // dereference the ms pointer, need parenthesis cuz dereferencing operator has lower precedence than dot
	// accessing the value to the pointer which is a myStruct object, then accessing the foo field value
	fmt.Println((*ns).foo)
	ns.foo = 40 // syntactic sugar of the line above, Go will interpret it as the line above
	fmt.Println(ns.foo)

	// Slice pointer
	mySlice := []int{1, 2, 3} // slice is a pointer to an underlying array
	cloneSlice := mySlice     // what's being copied is the pointer to the underlying array
	fmt.Println(mySlice, cloneSlice)
	mySlice[1] = 10
	fmt.Println(mySlice, cloneSlice)

	// Map has the same pointer behavior as the slice pointer
	// have a pointer to an underlying data, doesnt' contain the underlying data itself
	myMap := map[string]string{"one": "two", "three": "six"}
	cloneMap := myMap
	fmt.Println(myMap, cloneMap)
	myMap["one"] = "hundred"
	fmt.Println(myMap, cloneMap)
}

type myStruct struct {
	foo int
}
