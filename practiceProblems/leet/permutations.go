package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	answer := [][]int{}
	traversePerms(nums, []int{}, &answer)
	return answer
}

func traversePerms(remainingNums, current []int, answer *[][]int) {
	if len(remainingNums) == 1 {
		current = append(current, remainingNums[0])
		*answer = append(*answer, current)
	}
	for i := 0; i < len(remainingNums); i++ {
		nextRemaining := append([]int{}, remainingNums...)
		nextRemaining = append(nextRemaining[:i], nextRemaining[i+1:]...)
		nextCurr := append([]int{}, current...)
		nextCurr = append(nextCurr, remainingNums[i])
		traversePerms(nextRemaining, nextCurr, answer)
	}
}
