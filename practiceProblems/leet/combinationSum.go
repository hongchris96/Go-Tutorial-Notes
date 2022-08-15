package main

import (
	"fmt"
)

func main() {
	fmt.Println("Attempt 1")
	fmt.Println(combinationSum([]int{100, 200, 4, 12}, 400)) // hitting infinite loop
}

// DFS structure for [2,3,6,7] target 7:
// 													[]
// 							[2]                                           []
// 			[2,2]                  [2]								[3]				[]
// 	[2,2,2]       [2,2,3]		[2,3]		 [2]			  [3,3]		[3]		   [6]		[]
// [2,2,2,2]				[2,3,3] [2,3,6]		[2,6]	   [3,3,3]		 [3,6]	  [6,6]			[7]

func combinationSum(candidates []int, target int) [][]int {
	answer := [][]int{}
	// need to pass in slice by reference
	traverseCombs(candidates, []int{}, &answer, 0, target) // O(2^target): deepest dfs layer is the target
	return answer
}

func traverseCombs(candidates, current []int, answer *[][]int, total, target int) {
	if total == target {
		comb := append([]int{}, current...)
		*answer = append(*answer, comb)
		return
	} else if len(candidates) <= 0 || total > target {
		return
	}

	// includes the first candidate
	current = append(current, candidates[0])
	traverseCombs(candidates, current, answer, total+candidates[0], target)

	// no longer includes the first candidate
	current = append(current[:len(current)-1], current[len(current):]...)
	traverseCombs(candidates[1:], current, answer, total, target)
}
