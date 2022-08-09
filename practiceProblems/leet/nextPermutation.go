package main

import "fmt"

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
	// 125043
	// 125
	// 2134
	// 2143
	// 2314

	largestIdx := 0
	largestNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > largestNum {
			largestNum = nums[i]
			largestIdx = i
		}
	}

	// doable operation with right half
}
