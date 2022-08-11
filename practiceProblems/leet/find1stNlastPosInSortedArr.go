package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{1, 1, 1, 1, 1, 1, 2, 3, 4, 4, 5, 5, 5, 6, 7, 8, 8, 8, 8}, 8)) // 15, 18
}

func searchRange(nums []int, target int) []int {
	answer := []int{-1, -1}

	// binary search for first occurrence
	start := 0
	end := len(nums) - 1
	var i int
	for start <= end {
		i = (end + start) / 2
		if nums[i] == target {
			// set answer to this index
			answer[0], answer[1] = i, i
			break
		} else if nums[i] > target {
			end = i - 1
		} else {
			start = i + 1
		}
	}

	// if not found
	if answer[0] == -1 {
		return answer
	}

	if i != 0 && nums[i] == nums[i-1] {
		// binary search left and update left bound answer[0]
		answer[0] = bSearchExtra(nums[:i], target, -1)
	}
	if i != len(nums)-1 && nums[i] == nums[i+1] {
		// binary search right and update right bound answer[1]
		answer[1] = bSearchExtra(nums[i+1:], target, 1) + i + 1
	}
	return answer
}

func bSearchExtra(nums []int, target, direction int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		i := (end + start) / 2
		// found target value and neighbor value is not the same
		// left or right neighbor depending on the direction of this search
		if nums[i] == target && (i+direction > len(nums)-1 || i+direction < 0 || nums[i+direction] != target) {
			return i
		} else if direction == 1 { // searching for right bound, value can only be the same or larger
			// right neighbor value is still the same, go right
			if nums[i] == target {
				start = i + 1
			} else { // current value is larger than target, go left
				end = i - 1
			}
		} else if direction == -1 { // searching for left bound, value can only be the same or smaller
			// left neighbor value is still the same, go left
			if nums[i] == target {
				end = i - 1
			} else { // current value is smaller, go right
				start = i + 1
			}
		}
	}
	return -1
}
