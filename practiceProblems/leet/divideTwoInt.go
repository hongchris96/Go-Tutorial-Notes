package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(10, 3))
}

func divide(dividend int, divisor int) int {
	// updated bit operation solution from https://www.youtube.com/watch?v=htX69j1jf5U
	if math.MinInt32 == dividend && divisor == -1 {
		return math.MaxInt32
	}
	num1 := ab(dividend)
	num2 := ab(divisor)
	count := 0
	for num1-num2 >= 0 {
		x := 0 // 2 to the 0
		for num1-(num2<<1<<x) >= 0 {
			x++
		}
		count += 1 << x
		num1 -= num2 << x
	}
	if (dividend >= 0) == (divisor >= 0) {
		return -count
	}
	return count
}

func ab(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
