package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(19, 13)) // 86493225
}

func uniquePaths(m int, n int) int {
	cached := map[[2]int]int{}
	return traverse(0, 0, m-1, n-1, cached)
}

func traverse(row, col, destRow, destCol int, cached map[[2]int]int) int {
	if row == destRow && col == destCol {
		return 1
	}
	if val, ok := cached[[2]int{row, col}]; ok {
		return val
	}
	count := 0
	if row < destRow {
		count += traverse(row+1, col, destRow, destCol, cached)
	}
	if col < destCol {
		count += traverse(row, col+1, destRow, destCol, cached)
	}
	cached[[2]int{row, col}] = count
	return count
}
