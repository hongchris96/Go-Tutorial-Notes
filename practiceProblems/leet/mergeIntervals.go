package main

import (
	"fmt"
	"math"
	"sort"
)

type customSlice [][]int

func main() {
	fmt.Println(merge([][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}})) // [[1, 10]]
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))       // [[1,6],[8,10],[15,18]]
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))                          // [[1, 5]]
}

func (s customSlice) Len() int {
	return len(s)
}
func (s customSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s customSlice) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}

func merge(intervals [][]int) [][]int {
	sort.Sort(customSlice(intervals))
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
