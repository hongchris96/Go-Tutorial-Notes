package main

import (
	"fmt"
)

func main() {
	// Interface
	fmt.Println("Interface:")
	// w is holding a Writer, which is something that implements the Writer interface
	// don't know the concrete type of w
	var w Writer = ConsoleWriter{} // polymorphic behavior
	// but we know how to call the method on w because of the interface
	w.Write([]byte("Hellow Go!"))

	// can use any custom type to implement interfaces
	myInt := CustomInt(0)
	var inc Incrementer = &myInt // assign pointer to custom int to interface
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

	// interface type conversion (not too sure how it works)
	// var wc WriterCloser = Something()
	// bwc := wc.(*Something())

	// Empty Interface:
	// interface with no methods on it
	var myObj interface{} = []int{}
	fmt.Println(myObj)

	// Type switch
	var i interface{} = 0
	switch i.(type) {
	case int:
		fmt.Println("Integer")
	case string:
		fmt.Println("String")
	default:
		fmt.Println("Unknown")
	}

	// Note:
	// if implementing an interface using value type, the methods that implements the interface have to all have value receiveers.
	// If implementing an interface with a pointer, just have to have the methods there regardless of receiver type

	// Best Practice:
	// - Use many small interfaces instead of large monolithic ones
	// - don't export interfaces if you don't need to
	// - do export interfaces that will be used by package
	// - design func and methods to receive interfaces whenever possible

}

// interfaces describe behaviors not data
// stores method definition
// naming convention ends with er
type Writer interface {
	// method here accepts a slice of byte, return int and error
	Write([]byte) (int, error)
}

// implicitly implement inferface with a struct
type ConsoleWriter struct{}

// define method on ConsoleWriter called Write
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// interface implemented on custom type
type Incrementer interface {
	Increment() int
}

type CustomInt int

// method on pointer of custom int
func (i *CustomInt) Increment() int {
	*i++           // increment the derefenced value
	return int(*i) // return the dereference value converted to int
}

// embedding interfaces
type Righter interface {
	Write([]byte) (int, error)
}
type Closer interface {
	Close() error
}

// composed interface
type WriterCloser interface {
	Righter
	Closer
}
