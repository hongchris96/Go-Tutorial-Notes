package main

import (
	"fmt"
)

const enumConst = iota // iota is a counter to used when creating an enumerated constant

// each iota is scoped to their own constant block
const (
	enumA = iota
	enumB = iota
	enumC = iota
)

// if we don't assign value after the first one, compiler will try to infer the pattern of the assignments
// applying to the same formula of D to E and F
const (
	enumD = iota
	enumE
	enumF
)

func main() {
	// use Title case if exporting this const
	// const MyConst

	// constants has to be assignable at compile time -> primitives
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)

	// constants cannot be set to something that has be determined at run time
	// const myConst float64 = math.Sin(1.57) will throw an error

	// math operable between const and var if they are the same type
	var a int = 27
	fmt.Printf("%v, %T\n", myConst+a, myConst+a)

	// compiler has the ability to infer untyped constant data type
	// compiler will look at where the constant is being used and then implicitly convert type
	const b = 42
	fmt.Printf("%v, %T\n", b, b) // type is int
	var c int16 = 27
	fmt.Printf("%v, %T\n", b+c, b+c) // type is int16

	// --------------------------------------------------------------------------

	// Enumerated Constant (see above function main)
	fmt.Println("Enumerated Constants:")

	fmt.Printf("%v, %T\n", enumConst, enumConst) // type is int
	fmt.Printf("%v\n", enumA)
	fmt.Printf("%v\n", enumB)
	fmt.Printf("%v\n", enumC)
	fmt.Printf("%v\n", enumD)
	fmt.Printf("%v\n", enumE)
	fmt.Printf("%v\n", enumF)
}
