package main

import (
	"fmt"
)

func main() {
	// For loops
	fmt.Println("For Loops:")
	// for initializer; boolean statement; incrementer { task in each iteration }
	for i := 0; i < 5; i++ {
		// i is scoped within the for loop block
		fmt.Println(i)
	}

	// multiple incrementers
	// cannot have two statements in the initializer and incrementer:			for i = 0, j = 0;
	// can do both assignments together by combining them into one statement 	for i, j = 0, 0;
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	// can leave the initializer empty, can also leave the incrementer empty
	outerI := 0 // counter is now globally scoped
	// must have the first semicolon so that everything else is in the correct position
	for ; outerI < 5; outerI++ {
		fmt.Println(outerI)
	}
	fmt.Println("Outside the for loop:", outerI)

	// While loop: just a special case of for loop
	outerI = 0
	fmt.Println("While Loops:")
	for outerI < 5 { // syntactic sugar of for ; outerI < 5 ; {}
		fmt.Println(outerI)
		// moving the incrementer inside
		outerI++
	}
	// while true loop
	// infinite loop until break called inside
	for {
		fmt.Println(outerI)
		outerI++
		if outerI == 10 {
			break
		}
	}

	// break and continue
	// break: leaves the current loop
	// continue: leaves the current iteration and on to the next

	// breaking out of outer loop in a nested loop using labels
outerLoopLabel:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break outerLoopLabel // breaks out of inner loop and outer loop
			}
		}
	}

	// For range loop
	// special form of for loop to loop through a collection
	fmt.Println("For Range Loop:")
	collection := []int{1, 2, 3} // can be slice, array, map, string
	for k, v := range collection {
		fmt.Println(k, v)
		// slice and array:	k -> index, v -> element
		// map: 			k -> key, v -> value
		// string: 			k -> index, v -> char unicode
	}
	// for range initializer variable must be used, if you know one is not gonna be used, replace it with _
	for k, _ := range collection {
		fmt.Println(k)
		// slice and array:	k -> index, v -> element
		// map: 			k -> key, v -> value
		// string: 			k -> index, v -> char unicode
	}
}
