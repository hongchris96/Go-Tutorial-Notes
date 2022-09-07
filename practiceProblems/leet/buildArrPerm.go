package main

import "fmt"

func main() {
	// [nums[nums[0]], nums[nums[1]], nums[nums[2]], nums[nums[3]], nums[nums[4]], nums[nums[5]]]
	fmt.Println(buildArray([]int{0, 2, 1, 5, 3, 4})) // [0,1,2,4,5,3]
	fmt.Println(buildArray([]int{5, 0, 1, 2, 3, 4})) // [4,5,0,1,2,3]
}

func buildArray(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		nums[i] = nums[i] + (n * (nums[nums[i]] % n))
	}

	for i := 0; i < n; i++ {
		nums[i] = nums[i] / n
	}

	return nums
}
