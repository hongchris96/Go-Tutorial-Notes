package main

import "fmt"

func main() {
	fmt.Println("Zig Zag Convert")
	fmt.Println(convert("PAYPALISHIRING", 3) == "PAHNAPLSIIGYIR") // PAHNAPLSIIGYIR
	// P   A   H   N
	// A P L S I I G
	// Y   I   R

	fmt.Println(convert("PAYPALISHIRING", 4) == "PINALSIGYAHRPI") // PINALSIGYAHRPI
	// P     I    N
	// A   L S  I G
	// Y A   H R
	// P     I
}

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	// make 2D slices of numRows
	var picture []string
	for i := 0; i < numRows; i++ {
		picture = append(picture, "")
	}
	rowTracker := 0
	directionDown := false
	// loop through string, append character to correct slice in 2D slices
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		picture[rowTracker] += char

		if rowTracker == 0 || rowTracker == numRows-1 {
			directionDown = !directionDown
		}

		if directionDown {
			rowTracker += 1
		} else {
			rowTracker -= 1
		}
	}

	answer := ""
	for i := 0; i < len(picture); i++ {
		answer += picture[i]
	}

	return answer
}
