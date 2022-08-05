package main

import "fmt"

func main() {
	slice1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	slice2 := []int{1, 1, 2}
	fmt.Println(removeDuplicates(slice1))
	fmt.Println(slice1)
	fmt.Println(removeDuplicates(slice2))
	fmt.Println(slice2)
}

func removeDuplicates(nums []int) int {
	i := 0
	for i < len(nums) {
		if i > 0 {
			firstDup := i - 1
			lastDup := i
			for lastDup < len(nums) && nums[firstDup] == nums[lastDup] {
				lastDup++
			}
			nums = append(nums[:firstDup+1], nums[lastDup:]...)
			if lastDup != i {
				i--
			}
		}
		i++
	}
	return len(nums)
}
