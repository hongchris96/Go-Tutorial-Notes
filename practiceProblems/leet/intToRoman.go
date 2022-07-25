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
	roman := ""
	zeroes := 1
	for num > 0 {
		digit := num % 10
		targetNum := digit * zeroes
		numCreated := 0
		romanCreated := ""
		if digit == 4 {
			romanCreated = translation[zeroes] + translation[5*zeroes]
			numCreated += 4 * zeroes
		} else if digit == 9 {
			romanCreated = translation[zeroes] + translation[10*zeroes]
			numCreated += 9 * zeroes
		} else if digit >= 5 {
			romanCreated = translation[5*zeroes]
			numCreated += 5 * zeroes
		}
		for numCreated < targetNum {
			romanCreated += translation[zeroes]
			numCreated += zeroes
		}
		num /= 10
		zeroes *= 10
		roman = romanCreated + roman
	}
	return roman
}
