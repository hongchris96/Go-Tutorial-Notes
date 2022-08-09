package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{5, 4, 7, 5, 3, 2}
	nextPermutation(a)
	fmt.Println(a)
}

func nextPermutation(nums []int) {
	lastIdx := len(nums) - 1
	swapIdx1 := lastIdx - 1
	// iterate left until the swapIdx1 has a right neighbor that's greater
	for swapIdx1 >= 0 && nums[swapIdx1] >= nums[swapIdx1+1] {
		swapIdx1--
	}
	if swapIdx1 > -1 {
		swapIdx2 := lastIdx
		// iterate left until swapIdx2 is greater than swapIdx1
		for swapIdx2 > swapIdx1 && nums[swapIdx1] >= nums[swapIdx2] {
			swapIdx2--
		}
		// swap the numbers
		temp := nums[swapIdx2]
		nums[swapIdx2] = nums[swapIdx1]
		nums[swapIdx1] = temp
	}
	// sort the right half from the swapping point
	// (if the nums are in sorted descending order, sorts the whole slice)
	sort.Ints(nums[swapIdx1+1:])
}
