package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println()
}

func merge(intervals [][]int) [][]int {
	for i := 1; i < len(intervals); i++ {
		prevLower := float64(intervals[i-1][0])
		prevUpper := float64(intervals[i-1][1])
		lowerBound := float64(intervals[i][0])
		upperBound := float64(intervals[i][1])
		if (lowerBound <= prevUpper && upperBound >= prevUpper) ||
			(prevLower <= upperBound && prevUpper >= upperBound) {
			intervals[i-1][0] = int(math.Min(prevLower, lowerBound))
			intervals[i-1][1] = int(math.Max(prevUpper, upperBound))
			intervals = append(intervals[:i], intervals[i+1:]...)
			i--
		}
	}
	return intervals
}
