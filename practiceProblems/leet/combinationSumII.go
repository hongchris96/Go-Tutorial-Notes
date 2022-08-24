package main

import "fmt"

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

func combinationSum2(candidates []int, target int) [][]int {
	answer := [][]int{}
	uniqueCandidates := []int{}
	candidatesCount := map[int]int{}
	for _, num := range candidates {
		if _, ok := candidatesCount[num]; ok {
			candidatesCount[num] += 1
		} else {
			candidatesCount[num] = 1
			uniqueCandidates = append(uniqueCandidates, num)
		}
	}
	traverseCombs(uniqueCandidates, []int{}, &answer, 0, target, candidatesCount)
	return answer
}

func traverseCombs(candidates, current []int, answer *[][]int, total, target int, candCount map[int]int) {
	if total == target {
		comb := append([]int{}, current...)
		*answer = append(*answer, comb)
		return
	} else if len(candidates) <= 0 || total > target {
		return
	}

	// includes the first candidate
	current = append(current, candidates[0])
	if candCount[candidates[0]] > 0 {
		candCount[candidates[0]] -= 1
		traverseCombs(candidates, current, answer, total+candidates[0], target, candCount)
	}

	// no longer includes the first candidate
	current = append(current[:len(current)-1], current[len(current):]...)
	candCountCopy := mapCopy(candCount)
	traverseCombs(candidates[1:], current, answer, total, target, candCountCopy)
}

func mapCopy(mapOriginal map[int]int) map[int]int {
	newMap := map[int]int{}
	for k, v := range mapOriginal {
		newMap[k] = v
	}
	return newMap
}
