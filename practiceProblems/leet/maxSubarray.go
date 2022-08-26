package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(maxSubArray([]int{-2, -3, -1, -2, -4, -2}))        // -1
}

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currentSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if currentSum < 0 {
			currentSum = nums[i]
		} else {
			currentSum += nums[i]
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}
