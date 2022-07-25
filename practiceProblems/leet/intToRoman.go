package main

import "fmt"

func main() {
	fmt.Println(intToRoman(3729)) // MMMDCCXXIX
}

func intToRoman(num int) string {
	translation := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}
	romanTotal := ""
	zeroes := 1 // zeroes are 1 and multiple of 10s
	for num > 0 {
		digit := num % 10
		// building roman symbol for the current digit * zeroes
		targetNum := digit * zeroes
		numCreated := 0
		romanCreated := ""
		if digit == 4 || digit == 9 {
			// the subtraction display
			romanCreated = translation[zeroes] + translation[(digit+1)*zeroes]
			numCreated += digit * zeroes
		} else if digit >= 5 {
			romanCreated = translation[5*zeroes]
			numCreated += 5 * zeroes
		}
		// if number value created so far is still less than target number, add trailing symbols
		for numCreated < targetNum {
			romanCreated += translation[zeroes]
			numCreated += zeroes
		}
		num /= 10
		zeroes *= 10
		romanTotal = romanCreated + romanTotal
	}
	return romanTotal
}
