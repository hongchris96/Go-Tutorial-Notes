package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3}
	nextPermutation(a)
	fmt.Println(a)
}

func nextPermutation(nums []int) {
	// 1234
	// 1243
	// 1324
	// 1342
	// 1423
	// 1432

	// 12453
	// 125034
	// 125430
	// 125
	// 2134
	// 2143
	// 2314
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
	} else if largestIdx == len(nums)-2 { // still need to handle largestIdx at 0
		swapNsort(nums, largestIdx)
	} else {
		if decSorted(nums[largestIdx+1:]) {
			swapNsort(nums[largestIdx-1:])
		} else {
			findPeak(nums[largestIdx+1:])
		}
	}
}

func swapNsort(nums []int, peak int) {
	swapingNum := nums[len(nums)-1]
	swappable := false
	for i := peak - 1; i >= 0; i-- {
		if nums[i] < swapingNum {
			swappable = true
			break
		}
	}
	if swappable {
		i := len(nums) - 1
		for ; i >= 0; i-- {
			if nums[i] < swapingNum {
				temp := nums[i]
				nums[i] = swapingNum
				nums[i+1] = temp
				break
			}
			nums[i] = nums[i-1]
		}
		sort.Ints(nums[i+1:])
	} else {
		temp := nums[peak-1]
		nums[peak-1] = nums[peak]
		nums[peak] = temp
		sort.Ints(nums[peak:])
	}
}

func decSorted(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			return false
		}
	}
	return true
}
