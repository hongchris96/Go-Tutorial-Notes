package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) //[[-1,-1,2],[-1,0,1]]
	fmt.Println(threeSum([]int{0, 1, 1}))             // []
	fmt.Println(threeSum([]int{0, 0, 0}))             // [0,0,0]
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	uniqueTriplets := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 { // if positive no num after can sum to zero
			break
		}
		if i > 0 && nums[i] == nums[i-1] { // prevent duplicate
			continue
		}
		uniqueTriplets = append(uniqueTriplets, twoSum(nums[i+1:], nums[i])...)
	}
	return uniqueTriplets
}

func twoSum(nums []int, firstNum int) [][]int {
	triplets := [][]int{}
	leftI := 0
	rightI := len(nums) - 1

	for leftI < rightI {
		currentSum := firstNum + nums[leftI] + nums[rightI]
		if currentSum == 0 {
			triplets = append(triplets, []int{firstNum, nums[leftI], nums[rightI]})
			leftI++ // moving on to next num
			goLeftTilDifferent(leftI, &rightI, nums)
		} else if currentSum > 0 {
			// need to lower the bigger positive number
			goLeftTilDifferent(leftI, &rightI, nums)
		} else {
			// need to increase the negative number
			leftI++
			// duplicate removal from the left side was handled above in the outer function
		}
	}
	return triplets
}

func goLeftTilDifferent(head int, i *int, nums []int) {
	// remove duplicate from the right side
	*i--
	for head < *i && nums[*i] == nums[*i+1] {
		*i--
	}
}
