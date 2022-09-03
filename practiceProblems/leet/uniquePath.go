package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(19, 13)) // 86493225
}

func uniquePaths(m int, n int) int {
	return traverse(0, 0, m-1, n-1)
}

func traverse(row, col, destRow, destCol int) int {
	if row == destRow && col == destCol {
		return 1
	}
	count := 0
	if row < destRow {
		count += traverse(row+1, col, destRow, destCol)
	}
	if col < destCol {
		count += traverse(row, col+1, destRow, destCol)
	}
	return count
}
