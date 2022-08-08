package main

import "fmt"

func main() {
	fmt.Println(divide(10, 3))
}

func divide(dividend int, divisor int) int {
	num1 := ab(dividend)
	num2 := ab(divisor)
	count := 0
	for num1 >= num2 {
		num1 -= num2
		count++
	}
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
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
