package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(3)) // [[1,2,3],[8,9,4],[7,6,5]]
}

func generateMatrix(n int) [][]int {
	answer := make([][]int, n)
	for i := range answer {
		answer[i] = make([]int, n)
	}

	row := 0
	col := 0
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dirIdx := 0
	currentDir := dirs[dirIdx]
	i := 1

	for i <= n*n {
		answer[row][col] = i
		nextRow, nextCol := row+currentDir[0], col+currentDir[1]
		if inRange(nextRow, nextCol, n) && answer[nextRow][nextCol] == 0 {
			row, col = nextRow, nextCol
		} else {
			dirIdx++
			currentDir = dirs[dirIdx%4]
			row, col = row+currentDir[0], col+currentDir[1]
		}
		i++
	}
	return answer
}

func inRange(row, col, length int) bool {
	if row > -1 && col > -1 && row < length && col < length {
		return true
	}
	return false
}
