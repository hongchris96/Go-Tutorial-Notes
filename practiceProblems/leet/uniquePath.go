package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(19, 13))                                                // 86493225
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}})) // 2
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

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	row := len(obstacleGrid)
	col := len(obstacleGrid[0])
	cached := map[[2]int]int{}
	return traverse2(0, 0, row-1, col-1, cached, obstacleGrid)
}

func traverse2(row, col, destRow, destCol int, cached map[[2]int]int, grid [][]int) int {
	if grid[row][col] == 1 {
		return 0
	}
	if row == destRow && col == destCol {
		return 1
	}
	if val, ok := cached[[2]int{row, col}]; ok {
		return val
	}
	count := 0
	if row < destRow {
		count += traverse2(row+1, col, destRow, destCol, cached, grid)
	}
	if col < destCol {
		count += traverse2(row, col+1, destRow, destCol, cached, grid)
	}
	cached[[2]int{row, col}] = count
	return count
}
