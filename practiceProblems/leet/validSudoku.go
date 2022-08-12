package main

import (
	"fmt"
)

func main() {
	fmt.Println("valid sudoku")
}

func isValidSudoku(board [][]byte) bool {
	rowBuckets := [9][9]int{}
	colBuckets := [9][9]int{}
	boxBuckets := [9][9]int{}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			numByte := board[row][col]
			if numByte != 46 {
				num := int(numByte - 48)
				if rowBuckets[row][num-1] != num {
					rowBuckets[row][num-1] = num
				} else {
					return false
				}

				if colBuckets[col][num-1] != num {
					colBuckets[col][num-1] = num
				} else {
					return false
				}

				box := boxAssign(row, col)
				if boxBuckets[box][num-1] != num {
					boxBuckets[box][num-1] = num
				} else {
					return false
				}

			}
		}
	}
	return true
}

func boxAssign(row, col int) int {
	switch {
	case row < 3 && col < 3:
		return 0
	case row < 3 && col < 6:
		return 1
	case row < 3 && col < 9:
		return 2
	case row < 6 && col < 3:
		return 3
	case row < 6 && col < 6:
		return 4
	case row < 6 && col < 9:
		return 5
	case row < 9 && col < 3:
		return 6
	case row < 9 && col < 6:
		return 7
	default:
		return 8
	}
}
