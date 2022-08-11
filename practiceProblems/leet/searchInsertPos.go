package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3}, 4))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	fmt.Println(searchInsert([]int{1}, 1))
}

func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		leftBound := (start + end) / 2
		rightBound := leftBound + 1
		if rightBound > len(nums)-1 {
			if nums[leftBound] < target {
				return leftBound + 1
			} else {
				return leftBound
			}
		} else if leftBound == 0 {
			if nums[leftBound] >= target {
				return leftBound
			} else if nums[rightBound] >= target {
				return rightBound
			} else {
				return rightBound + 1
			}
		}
		if target <= nums[rightBound] && target > nums[leftBound] {
			return rightBound
		} else if target <= nums[leftBound] {
			end = leftBound - 1
		} else if target > nums[rightBound] {
			start = rightBound
		}
	}

	return -1
}
