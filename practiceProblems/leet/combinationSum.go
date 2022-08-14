package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Attempt 1")
	fmt.Println(combinationSum([]int{100, 200, 4, 12}, 400)) // hitting infinite loop
}

func combinationSum(candidates []int, target int) [][]int {
	if target == 0 {
		return [][]int{[]int{}}
	}

	// assuming candidates are sorted in ascending order
	allCombs := [][]int{}
	for i := len(candidates) - 1; i >= 0; i-- {
		if candidates[i] <= target {
			remaining := target - candidates[i]
			remainingComb := combinationSum(candidates, remaining)
			for _, comb := range remainingComb {
				candidateComb := append([]int{candidates[i]}, comb...)
				allCombs = append(allCombs, candidateComb)
			}
		}
	}

	checked := [][]int{}
	for i := 0; i < len(allCombs); i++ {
		dupFound := false
		for _, checkedComb := range checked {
			if combDup(checkedComb, allCombs[i]) {
				dupFound = true
				break
			}
		}
		if dupFound {
			copy(allCombs[i:], allCombs[i+1:])
			allCombs = allCombs[:len(allCombs)-1]
			i--
		} else {
			checked = append(checked, allCombs[i])
		}
	}
	return allCombs
}

func combDup(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	arr1 = append([]int{}, arr1...)
	arr2 = append([]int{}, arr2...)
	sort.Ints(arr1)
	sort.Ints(arr2)
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
