package main

import "fmt"

func main() {
	fmt.Println(myPow(2.1000, 3))  // 9.26100
	fmt.Println(myPow(2.0000, -2)) // 0.25000
}

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	ans := 1.0
	for n > 0 {
		lastBit := n & 1
		if lastBit > 0 {
			ans = ans * x
		}
		x *= x
		// Right shift
		n = n >> 1
	}
	return ans
}
