package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})) // [1,2,3,6,9,8,7,4,5]
}

func spiralOrder(matrix [][]int) []int {
	visited := map[[2]int]bool{}
	answer := []int{}
	row, col := 0, 0
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dirOption := 0
	for {
		answer = append(answer, matrix[row][col])
		visited[[2]int{row, col}] = true
		if len(answer) == len(matrix)*len(matrix[0]) {
			return answer
		}
		for {
			dir := dirs[dirOption%4]
			nextRow, nextCol := row+dir[0], col+dir[1]
			_, exist := visited[[2]int{nextRow, nextCol}]
			if inRange(nextRow, nextCol, len(matrix), len(matrix[0])) && !exist {
				row, col = nextRow, nextCol
				break
			}
			dirOption++
		}
	}
	return answer
}

func inRange(row, col, rowLimit, colLimit int) bool {
	if row > -1 &&
		col > -1 &&
		row < rowLimit &&
		col < colLimit {
		return true
	}
	return false
}
