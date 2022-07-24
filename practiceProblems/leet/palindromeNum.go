package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	// digits := []int{}
	// for x > 0 {
	// 	digit := x % 10
	// 	digits = append(digits, digit)
	// 	x /= 10
	// }

	// for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
	// 	if digits[i] != digits[j] {
	// 		return false
	// 	}
	// }

	// return true

	reverseNum := 0
	forwardNum := x

	for x > 0 {
		digit := x % 10
		reverseNum = reverseNum*10 + digit
		x /= 10
	}

	return reverseNum == forwardNum
}
