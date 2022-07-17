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
}
