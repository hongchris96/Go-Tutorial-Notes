package main

import "fmt"

func main() {
	fmt.Println()
}

func rotate(matrix [][]int) {
	// 00 01 02 03 -> 03 13 23 33 (row 0, second num 3)
	// 10 11 12 13 -> 02 12 22 32 (row 1, second num 2)
	// 20 21 22 23 -> 01 11 21 31 (row 2, second num len-row-1)
	// second num takes the place of first num
	// (len - first num) takes the place of second num

	newPosition := map[[2]int]int{}
	matrixLen := len(matrix)
	// assign new position for each val
	for row := 0; row < matrixLen; row++ {
		for col := 0; col < matrixLen; col++ {
			val := matrix[row][col]
			newPosition[[2]int{col, matrixLen - row - 1}] = val
		}
	}
	for position, val := range newPosition {
		matrix[position[0]][position[1]] = val
	}
}
