package main

import (
	"fmt"
)

func main() {
	fmt.Println(jump([]int{5, 6, 4, 4, 6, 9, 4, 4, 7, 4, 4, 8, 2, 6, 8, 1, 5, 9, 6, 5, 2, 7, 9, 7, 9, 6, 9, 4, 1, 6, 8, 8, 4, 4, 2, 0, 3, 8, 5}))
}

func jump(nums []int) int {
	steps, jumpStart, currentFurthest := 0, 0, 0

	for currentFurthest < len(nums)-1 {
		furthestJump := 0
		// picking jump pads that can jump further
		for i := jumpStart; i < currentFurthest+1; i++ {
			if i+nums[i] > furthestJump {
				furthestJump = i + nums[i]
			}
		}
		jumpStart = currentFurthest + 1
		currentFurthest = furthestJump
		steps += 1
	}

	return steps
}
