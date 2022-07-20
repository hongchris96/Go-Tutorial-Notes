package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	nums2 := []int{3, 2, 4}
	target2 := 6

	nums3 := []int{3, 3}
	target3 := 6

	fmt.Println(twoSum(nums, target))   // [0, 1]
	fmt.Println(twoSum(nums2, target2)) // [1, 2]
	fmt.Println(twoSum(nums3, target3)) // [0, 1]
}

func twoSum(nums []int, target int) []int {
	// for i, num := range nums {
	//     pairNum := target - num
	//     for j, num2 := range nums[i+1:] {
	//         if num2 == pairNum {
	//             return []int{i, i + j + 1}
	//         }
	//     }
	// }

	// for i := 0; i < len(nums)-1; i++ {
	//     pairNum := target - nums[i]
	//     for j := i + 1; j < len(nums); j++ {
	//         if nums[j] == pairNum {
	//             return []int{i, j}
	//         }
	//     }
	// }

	checkedNums := map[int]int{}
	for i, num := range nums {
		if idx, ok := checkedNums[target-num]; ok {
			return []int{i, idx}
		} else {
			checkedNums[num] = i
		}
	}
	return []int{}
}
