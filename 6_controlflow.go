package main

import (
	"fmt"
	"math"
)

func main() {

	// IF statement
	fmt.Println("IF:")

	// style 1: if conditional {}
	if true {
		fmt.Println("This is the first conditional")
	}

	// style 2: initializer syntax
	// sample data
	sampleMap := map[string]int{
		"Dio":   1234521,
		"Jojo":  3523521,
		"Wagon": 1234323,
		"Smoky": 4538903,
		"Lisa":  1122323,
	}
	// if initializer; conditional { variable in initializer is accessible within this scope }
	if value, exist := sampleMap["Dio"]; exist {
		fmt.Println("Dio's number is", value)
	}

	// comparison operator
	// <, >, <=, >=, ==, !=

	// boolean operator
	// || or, && and, ! not
	// shortcircuiting applies

	// else statements
	number := 50
	guess := 30
	if guess < 1 {
		fmt.Println("Must be greater than 1")
	} else if guess > 100 {
		fmt.Println("Must be less than 100")
	} else {
		if guess == number {
			fmt.Println("You are correct")
		} else {
			fmt.Println("You are within the ballpark")
		}
	}

	// be aware of floating points and decimals
	// go uses floats here, floats are approximation of decimal values (not exact)
	myNum := 0.123
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("Square rooting then go up to the power of 2 gives the same")
	} else {
		fmt.Println("Square rooting then go up to the power of 2 gives different val")
	}
	// can only solve it by using margin of error
	// divide the two number will be close to 1, subtract 1, and check if margin of error is within certain threshold
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
		fmt.Println("The are close enough to be the same")
	}

	// --------------------------------------------------------------------------

	// Switch
	fmt.Println("Switch:")
	// switch case that checks the value of switchNum
	switchNum := 5
	// switchNum here is called a tag
	switch switchNum {
	case 1:
		fmt.Println("case one")
	case 2:
		fmt.Println("case two")
	// can combine cases into one, but each case value here must be unique and not overlap with other cases
	// e.g. cannot put another 2 in this group, 5 cannot be in other cases
	case 5, 6, 7:
		fmt.Println("case five six or seven")
	default:
		fmt.Println("default case")
	}

	// initializer syntax: switch initializer; tag(value that we're checking) {}
	switch i := 2 + 3; i {
	case 5:
		fmt.Println("i is 5")
	default:
		fmt.Println("default case")
	}

	// tagless syntax
	outerVar := 10
	switch {
	// by having no tags, each case can have a full comparison operation
	case outerVar <= 10:
		fmt.Println("less than or equal to 10")
		// notice there's no break keyword for each case, because they are implied in Go switch cases
		// but if you want it fall through like switch cases in other languages, use fallthrough keyword
		fallthrough
	case outerVar <= 20:
		fmt.Println("less than or equal to 20")
		fmt.Println("Can have multiple operations in a case block")
		break // break early to skip the operation below
		fmt.Println("Skip this statement")
	default:
		fmt.Println("greater than 20")
	}

	// Type switch: checks type of the value in the interface
	var x interface{} = [3]int{}
	switch x.(type) {
	case int:
		fmt.Println("x is integer")
	case float64:
		fmt.Println("x is float64")
	case string:
		fmt.Println("x is string")
	case [2]int:
		// not all arrays are equivalent
		// same array types must have the same length and same element type
		fmt.Println("x is a length 2 array")
	case [3]int:
		fmt.Println("x is a length 3 array")
	default:
		fmt.Println("x is something else")
	}

}
