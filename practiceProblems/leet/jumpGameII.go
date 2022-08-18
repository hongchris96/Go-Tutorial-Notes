package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(jump([]int{5, 6, 4, 4, 6, 9, 4, 4, 7, 4, 4, 8, 2, 6, 8, 1, 5, 9, 6, 5, 2, 7, 9, 7, 9, 6, 9, 4, 1, 6, 8, 8, 4, 4, 2, 0, 3, 8, 5}))
}

func jump(nums []int) int {
	return jumping(0, 0, nums)
}

func jumping(steps, currentLoc int, remainingNums []int) int {
	fmt.Println("steps", steps)
	fmt.Println("loc", currentLoc)
	stepsToEnd := math.MaxInt
	if currentLoc == len(remainingNums)-1 {
		return steps
	}
	for i := remainingNums[currentLoc]; i > 0; i-- {
		if (currentLoc+i < len(remainingNums)-1 && remainingNums[currentLoc+i] != 0) || currentLoc+i == len(remainingNums)-1 {
			numSteps := jumping(steps+1, currentLoc+i, remainingNums)
			if stepsToEnd > numSteps {
				stepsToEnd = numSteps
			}
		}
	}
	return stepsToEnd
}
