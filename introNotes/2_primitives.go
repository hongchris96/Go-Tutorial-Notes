package main

import (
	"fmt"
)

func main() {
	// Booleans
	fmt.Println("Booleans:")
	// explicit declaration
	var n bool = false
	fmt.Printf("%v, %T\n", n, n)

	// implicit
	m := 1 == 1
	fmt.Printf("%v, %T\n", m, m)

	// everytime you initialize a variable, default value is 0, and 0 is false
	var k int
	fmt.Printf("%v, %T\n", k, k)
	var t bool
	fmt.Printf("%v, %T\n", t, t)

	// --------------------------------------------------------------------------

	// Integers
	fmt.Println("Integers:")
	// int8		-128 to 127
	// int16	-32,768 to 32,767
	// int32	-2 billion to 2 billion (approx)
	// int64	-9 trillion to 9 trillion (approx)

	// uint8	0 to 255
	// uint16	0 to 65,535
	// uint32	0 to 4 billion (approx)

	// math operations must be done on same types
	a := 1
	var b int8 = 2 // will error when directly adding to a
	// therefore must to type conversion first
	fmt.Println(a + int(b))

	// bit operator
	fmt.Println("Bit Operation:")
	c := 10             // 1010
	d := 3              // 0011
	fmt.Println(c & d)  // and	: 0010	(both true)
	fmt.Println(c | d)  // or 	: 1011	(one true)
	fmt.Println(c ^ d)  // xor 	: 1001	(either one true but not both)
	fmt.Println(c &^ d) // andNot: 0100 (both false)

	// bit shift
	fmt.Println("Bit Shift:")
	e := 10             // 1010
	fmt.Println(e << 3) // 1010000 shifting left meaning binary move to the left by adding 000 to the end
	fmt.Println(e >> 3) // 0001 shifting right meaning binary move to right by adding 000 to the front

	// --------------------------------------------------------------------------

	// Floats
	fmt.Println("Floats:")
	// float32	-1.18E-38 to 3.4E38
	// float64	-2.23E-308 to 1.8E308

	// 3 ways to declare floats
	f := 3.14
	f = 13.7e72
	f = 2.1E14
	fmt.Printf("%v, %T\n", f, f)

	// operations not including remainder, bit op, bit shift
	g := 10.2
	h := 3.7
	fmt.Println(g + h)
	fmt.Println(g - h)
	fmt.Println(g * h)
	fmt.Println(g / h)

	// --------------------------------------------------------------------------

	// Complex Number
	fmt.Println("Complex Number:")
	// imaginary numbers literal syntax
	var i complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", i, i)
	// real() and imag() breaks the complex number apart
	// and return float32 for complex64, float64 for complex128
	fmt.Printf("%v, %T\n", real(i), real(i))
	fmt.Printf("%v, %T\n", imag(i), imag(i))

	// creating complex number using complex(realNumber, imagNumber)
	// the args are float64 if the complex is 128
	var j complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", j, j)

	// --------------------------------------------------------------------------

	// String
	fmt.Println("String:")
	var o string = "this is a string"
	// string are aliases for bytes
	fmt.Printf("%v, %T\n", o[2], o[2])
	// converting byte back to string
	fmt.Printf("%v, %T\n", string(o[2]), o[2])

	// strings are immutable
	// o[2] = "u" will error due to different types (uint8, string) and cannot mutate string

	// string concat
	p := " add on"
	fmt.Println(o + p)

	// converting strings to collection of bytes
	q := []byte(p)
	fmt.Printf("%v, %T\n", q, q)

	// --------------------------------------------------------------------------

	// Rune
	// single quote with one character
	// true type alias of int32
	fmt.Println("Rune:")
	r := 'a'
	fmt.Printf("%v, %T\n", r, r)
}
