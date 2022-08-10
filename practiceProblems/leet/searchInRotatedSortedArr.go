package main

import "fmt"

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	//						 0, 1, 2, 3, 4, 5, 6
	// fmt.Println(search([]int{2, 3, 4, 5, 6, 7, 1}, 2))
	//							 	 0, 1, 2, 3, 4, 5, 6
}

func search(nums []int, target int) int {
	leftBound, rightBound := findBoundaries(nums)
	start := leftBound
	end := rightBound

	for start != end {
		midIdx := (end + start) / 2
		if end < start {
			midIdx = ((end + start + len(nums)) / 2) % len(nums)
		}
		if nums[midIdx] > target {
			end = (midIdx - 1 + len(nums)) % len(nums)
		} else if nums[midIdx] < target {
			start = (midIdx + 1) % len(nums)
		} else {
			return midIdx
		}
	}
	if nums[start] == target {
		return start
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
