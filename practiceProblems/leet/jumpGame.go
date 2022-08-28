package main

import "fmt"

func main() {
	fmt.Println()
}

func canJump(nums []int) bool {
	if nums[0] >= len(nums)-1 {
		return true
	}
	one := 0
	two := 0

	for two < len(nums)-1 {
		if nums[two] == 0 {
			return false
		}
		if two == 0 {
			two = nums[0]
		}
		i := one
		currentTwo := two
		for ; i <= currentTwo; i++ {
			if i+nums[i] > two && i+nums[i] < len(nums)-1 && nums[i+nums[i]] != 0 {
				two = i + nums[i]
			} else if i+nums[i] >= len(nums)-1 {
				return true
			}
		}
		one = i
		if currentTwo == two {
			return false
		}
	}
	return true
}
