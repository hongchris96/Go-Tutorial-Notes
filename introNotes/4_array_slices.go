package main

import (
	"fmt"
)

func main() {
	// Array
	// fixed size to be known at compile time

	fmt.Println("Arrays: [...]int{1, 2, 3}")
	// declaring an array: [size of array]type{values of array}
	// contiguous in memory, works faster than having 3 separate variables
	myArrayLength3 := [3]int{1, 2, 3} // literal syntax
	fmt.Printf("Numbers: %v\n", myArrayLength3)
	// use [...] size to create an array just big enough to hold the values in the literal syntax
	myArrayLengthEnough := [...]int{1, 2, 3, 4}
	fmt.Printf("Numbers: %v\n", myArrayLengthEnough)

	// declaring empty array of fixed size
	var emptyArray [3]string
	fmt.Printf("Peoples: %v\n", emptyArray)
	emptyArray[0] = "Dio"
	emptyArray[2] = "Wagon"
	emptyArray[1] = "Jojo"
	fmt.Printf("Peoples: %v\n", emptyArray)
	fmt.Printf("People #2: %v\n", emptyArray[1])
	// len()
	fmt.Printf("People's length: %v\n", len(emptyArray))

	// 2D array
	var twoDee [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{1, 0, 1}}
	// same as
	// var twoDee [3][3]int
	// twoDee[0] = [3]int{1, 0, 0}
	// twoDee[2] = [3]int{0, 1, 0}
	// twoDee[1] = [3]int{0, 0, 1}
	fmt.Println(twoDee)

	// Copying arrays
	// when assigning a variable with the value of array to another variable,
	// it's creating a copy. (new variable does not point to reference of original)
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	// use pointers to point to the original array
	c := &a
	c[1] = 5
	fmt.Println(a)
	fmt.Println(c)

	// --------------------------------------------------------------------------

	// Slice
	// unknown length, a projection of underlying array
	fmt.Println("Slices: []int{1, 2, 3}")

	// copying slices: slices are naturally reference type
	sliceA := []int{1, 2, 3}
	// don't need &
	sliceB := sliceA
	sliceB[1] = 5
	fmt.Println(sliceA)
	fmt.Println(sliceB)
	// len()
	fmt.Printf("Slices length: %v\n", len(sliceA))
	// number of elements in the slice doesn't necessarily match the size of the array backbone
	fmt.Printf("Slices capacity: %v\n", cap(sliceA))

	// Array and Slice slicing operations
	sliceC := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// the four slices below are all pointing to the sliceC's underlying array
	// following operations will also work if sliceC is an array [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	allEle := sliceC[:]
	eleStartAt := sliceC[3:]  // ele starting including index 3 to the end
	eleEndAt := sliceC[:6]    // ele from start up to but not including index 6
	eleBetween := sliceC[3:6] // ele of index 3,4,5
	sliceC[5] = 100           // all of the sliced operation here are pointing to the same underlying array
	fmt.Println(allEle)
	fmt.Println(eleStartAt)
	fmt.Println(eleEndAt)
	fmt.Println(eleBetween)

	// creating slice using make(type, length of slice, optional length of underlying array)
	makeSlice := make([]int, 3, 100)
	fmt.Println(makeSlice) // initializes elements to 0
	fmt.Printf("Slices length: %v\n", len(makeSlice))
	fmt.Printf("Slices capacity: %v\n", cap(makeSlice))

	// appending to slice
	pendingSlice := []int{}
	fmt.Println(pendingSlice)
	fmt.Printf("Slices length: %v\n", len(pendingSlice))
	fmt.Printf("Slices capacity: %v\n", cap(pendingSlice))
	pendingSlice = append(pendingSlice, 1, 2, 3) // everything after first argument are appended elements
	fmt.Println(pendingSlice)
	fmt.Printf("Slices length: %v\n", len(pendingSlice))
	fmt.Printf("Slices capacity: %v\n", cap(pendingSlice))
	// appending to slice to results in size greater than capacity will creates a new underlying array
	// setting capacity through make([]int, 3, 100) allows you to use the same underlying array until 100+ elements

	// spread operator
	pendingSlice = append(pendingSlice, []int{4, 5}...)
	fmt.Println(pendingSlice)
	fmt.Printf("Slices length: %v\n", len(pendingSlice))
	fmt.Printf("Slices capacity: %v\n", cap(pendingSlice))

	// testing e.g.
	// newSlice here points to the same underlying array as pendingSlice
	newSlice := append(pendingSlice, 10)
	// mutating newSlice also mutates pendingSlice
	newSlice[3] = 20
	fmt.Println(newSlice, pendingSlice)
	fmt.Printf("Slices length: %v\n", len(newSlice))
	fmt.Printf("Slices capacity: %v\n", cap(newSlice))
	// newLargeSlice here has a different underlying array than pendingSlice
	// because the appending resulted in exceeding the capacity, causing a new underlying array to be made
	newLargeSlice := append(pendingSlice, 10, 10, 10, 10)
	// mutating newLargeSlice doesn't mutate pendingSlice
	newLargeSlice[3] = 40
	fmt.Println(newLargeSlice, pendingSlice)
	fmt.Printf("Slices length: %v\n", len(newLargeSlice))
	fmt.Printf("Slices capacity: %v\n", cap(newLargeSlice))
}
