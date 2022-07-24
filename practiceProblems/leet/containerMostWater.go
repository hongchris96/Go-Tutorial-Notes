package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
}

func maxArea(walls []int) int {
	leftWall := 0
	rightWall := len(walls) - 1
	globalMax := 0

	for leftWall < rightWall {
		leftHeight := walls[leftWall]
		rightHeight := walls[rightWall]
		currentArea := (rightWall - leftWall) * smaller(leftHeight, rightHeight)
		if currentArea > globalMax {
			globalMax = currentArea
		}
		if leftHeight > rightHeight {
			rightWall--
		} else {
			leftWall++
		}
	}

	return globalMax
}

func smaller(num1, num2 int) int {
	if num1 > num2 {
		return num2
	}
	return num1
}
