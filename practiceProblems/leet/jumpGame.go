package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))       // true
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))       // false
	fmt.Println(canJump([]int{1, 1, 2, 2, 0, 1, 1})) // true
}

func canJump(nums []int) bool {
	reachable := 0.0
	for i := 0; i < len(nums); i++ {
		// since i will keep moving forward
		// if the previous nums[i] is 0, reachable will stay in place making it less than i
		if reachable < float64(i) {
			return false
		}
		// update reachable based on current step
		reachable = math.Max(reachable, float64(nums[i]+i))
	}
	return true
}
