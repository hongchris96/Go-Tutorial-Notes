package main

import "fmt"

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{2, 3, 4, 5, 6, 7, 1}, 2))
}

func search(nums []int, target int) int {
	leftBound, rightBound := findBoundaries(nums)
	newArr := []int{}
	adjustment := 0
	if leftBound > rightBound {
		adjustment = leftBound
		newArr = append(newArr, nums[leftBound:]...)
		newArr = append(newArr, nums[:rightBound+1]...)
	} else {
		newArr = nums
	}
	start := 0
	end := len(nums) - 1

	for start <= end {
		midIdx := (start + end) / 2
		if newArr[midIdx] == target {
			return (midIdx + adjustment) % len(nums)
		} else if newArr[midIdx] > target {
			end = midIdx - 1
		} else {
			start = midIdx + 1
		}
	}
	return -1

}

func findBoundaries(nums []int) (int, int) {
	start := 0
	end := len(nums) - 1

	for start <= end {
		midIdx := (end + start) / 2
		rightIdx := (midIdx + 1) % len(nums)
		if nums[midIdx] > nums[rightIdx] {
			return rightIdx, midIdx
		}

		// max num is in the left half
		if nums[start] > nums[midIdx] {
			end = midIdx - 1
		} else { // max num in the right half
			start = midIdx + 1
		}
	}
	return 0, 0
}
