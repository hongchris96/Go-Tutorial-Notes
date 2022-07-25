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
	tens := 1
	for num > 0 {
		digit := num % 10
		targetNum := digit * tens
		numCreated := 0
		romanCreated := ""
		if digit == 4 {
			romanCreated = translation[tens] + translation[5*tens]
			numCreated += 4 * tens
		} else if digit == 9 {
			romanCreated = translation[tens] + translation[10*tens]
			numCreated += 9 * tens
		} else if digit >= 5 {
			romanCreated = translation[5*tens]
			numCreated += 5 * tens
		}
		for numCreated < targetNum {
			romanCreated += translation[tens]
			numCreated += tens
		}
		num /= 10
		tens *= 10
		roman = romanCreated + roman
	}
	return roman
}
