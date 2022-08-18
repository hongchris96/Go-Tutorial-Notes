package main

import (
	"fmt"
)

func main() {
	fmt.Println(jump([]int{5, 6, 4, 4, 6, 9, 4, 4, 7, 4, 4, 8, 2, 6, 8, 1, 5, 9, 6, 5, 2, 7, 9, 7, 9, 6, 9, 4, 1, 6, 8, 8, 4, 4, 2, 0, 3, 8, 5}))
}

func jump(nums []int) int {
	steps, _ := jumping(0, 0, nums)
	return steps
}

func jumping(steps, currentLoc int, remainingNums []int) (int, bool) {
	fmt.Println("steps", steps)
	fmt.Println("loc", currentLoc)
	stepsToEnd := []int{}
	if currentLoc == len(remainingNums)-1 {
		return steps, true
	}
	for i := remainingNums[currentLoc]; i > 0; i-- {
		if currentLoc+i < len(remainingNums) {
			numSteps, path := jumping(steps+1, currentLoc+i, remainingNums)
			if path {
				if len(stepsToEnd) == 0 {
					stepsToEnd = append(stepsToEnd, numSteps)
				} else if stepsToEnd[0] > numSteps {
					stepsToEnd[0] = numSteps
				}
			}
		}
	}
	if len(stepsToEnd) > 0 {
		return stepsToEnd[0], true
	}
	return -1, false
}
