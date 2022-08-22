package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPowerOfFour(16))
	fmt.Println(isPowerOfFour(8))
	fmt.Println(isPowerOfFour(1))
}

func isPowerOfFour(n int) bool {
	bit := strconv.FormatInt(int64(n), 2)
	if len(bit)%2 == 0 {
		return false
	}
	// I don't know how to bit manipulate
	for i, rune := range bit {
		if i == 0 && rune-48 != 1 {
			return false
		}
		if i > 0 && rune-48 != 0 {
			return false
		}
	}
	return true
}
