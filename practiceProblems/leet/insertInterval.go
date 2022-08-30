package main

import "fmt"

func main() {
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8})) // [[1,2],[3,10],[12,16]]
}

func insert(intervals [][]int, newInterval []int) [][]int {
	results := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		if newInterval[1] < intervals[i][0] {
			results = append(results, newInterval)
			results = append(results, intervals[i:]...)
			return results
		} else if newInterval[0] > intervals[i][1] {
			results = append(results, intervals[i])
		} else {
			newInterval[0] = min(newInterval[0], intervals[i][0])
			newInterval[1] = max(newInterval[1], intervals[i][1])
		}
	}
	results = append(results, newInterval)
	return results
}

func max(int1, int2 int) int {
	if int1 > int2 {
		return int1
	}
	return int2
}

func min(int1, int2 int) int {
	if int1 < int2 {
		return int1
	}
	return int2
}
