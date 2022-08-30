package main

import "fmt"

func main() {
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8})) // [[1,2],[3,10],[12,16]]
}

func insert(intervals [][]int, newInterval []int) [][]int {
	merged := false
	for i := 0; i < len(intervals); i++ {
		if !merged {
			if mergable(intervals[i], newInterval) {
				intervals[i][0] = min(intervals[i][0], newInterval[0])
				intervals[i][1] = max(intervals[i][1], newInterval[1])
				merged = true
			} else if insertable(intervals[i][0], newInterval) {
				intervals = append(intervals[:i+1], intervals[i:]...)
				intervals[i] = newInterval
				merged = true
			}
			continue
		}
		if merged && intervals[i][0] <= intervals[i-1][1] {
			intervals[i-1][1] = max(intervals[i-1][1], intervals[i][1])
			intervals = append(intervals[:i], intervals[i+1:]...)
			i--
		}
	}
	if !merged {
		intervals = append(intervals, newInterval)
	}
	return intervals
}

func insertable(lower int, interval []int) bool {
	if interval[1] < lower {
		return true
	}
	return false
}

func mergable(interval1, interval2 []int) bool {
	if interval2[0] <= interval1[1] && interval2[1] >= interval1[1] {
		return true
	} else if interval1[0] <= interval2[1] && interval1[1] >= interval2[1] {
		return true
	} else {
		return false
	}
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
