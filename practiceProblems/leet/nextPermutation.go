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
	if len(nums) > 1 {
		if decSorted(nums) {
			sort.Ints(nums)
		} else {
			findPeak(nums)
		}
	}
}

func findPeak(nums []int) {
	largestIdx := 0
	largestNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > largestNum {
			largestNum = nums[i]
			largestIdx = i
		}
	}
	if largestIdx == len(nums)-1 {
		temp := nums[largestIdx-1]
		nums[largestIdx-1] = largestNum
		nums[largestIdx] = temp
	} else if largestIdx == len(nums)-2 {
		swapNsort(nums, largestIdx)
	} else {
		if decSorted(nums[largestIdx:]) {
			swapNsort(nums[largestIdx-1:], 1)
		} else {
			findPeak(nums[largestIdx+1:])
		}
	}
}

func swapNsort(nums []int, peak int) {
	swapingNums := append([]int{}, nums[peak:]...)
	sort.Ints(swapingNums)
	i := 0
	swapingNum := swapingNums[i]
	swappable := false
	for !swappable {
		if nums[peak-1] < swapingNum {
			swappable = true
		}
		if !swappable {
			i++
			swapingNum = swapingNums[i]
		}
	}
	swapingNumIdx := -1
	i = len(nums) - 1
	for ; i >= 0; i-- {
		if swapingNumIdx == -1 {
			if swapingNum == nums[i] {
				swapingNumIdx = i
			}
			continue
		}
		if i < peak && nums[i] < swapingNum {
			temp := nums[i]
			nums[i] = swapingNum
			nums[swapingNumIdx] = temp
			break
		}
	}
	sort.Ints(nums[peak:])
}

func decSorted(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			return false
		}
	}
	return true
}
