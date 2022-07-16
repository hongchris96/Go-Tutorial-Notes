// Basic Struction

package main // package name

// all imports
import (
	"fmt" // needed for printing stuff
	"strconv"
)

// main function
func main() {
	fmt.Println("Sample raw number:", 42)

	// declaring variable with var
	// variable name of i
	// variable type of int
	var i int
	i = 42
	// can also be done in one line var i int = 42
	fmt.Println("Sample variable:", i)
	// {declared variables MUST be used or else compiling error}

	// go can figure out type, shorter declaration and assignment format here
	j := 99
	fmt.Println("Sample variable 2:", j)

	// printing the format of the variable: v is value, T is type
	fmt.Printf("%v, %T\n", j, j)

	// converting type by using typename(variable)
	var k float32
	k = float32(j)
	fmt.Println(("Type conversion to float:"))
	fmt.Printf("%v, %T\n", k, k)

	// when directly converting int to string, go finds the character with unicode based on int
	var l string
	l = string(j)
	fmt.Println(("Type conversion to string:"))
	fmt.Printf("%v, %T\n", l, l)

	// strconv package is needed to convert int 42 to string 42
	// Itoa converts Integer to askii value
	var m string
	m = strconv.Itoa(j)
	fmt.Println(("Type conversion to string using strconv:"))
	fmt.Printf("%v, %T\n", m, m)
}
