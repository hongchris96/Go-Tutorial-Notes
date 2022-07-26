package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"))     // 3
	fmt.Println(romanToInt("LVIII"))   // 58
	fmt.Println(romanToInt("MCMXCIV")) //1994
}

func romanToInt(s string) int {
	number := 0
	for i, bite := range s {
		letter := string(bite)
		switch letter {
		case "I":
			if i < len(s)-1 && (string(s[i+1]) == "V" || string(s[i+1]) == "X") {
				number -= 1
			} else {
				number += 1
			}
		case "V":
			number += 5
		case "X":
			if i < len(s)-1 && (string(s[i+1]) == "L" || string(s[i+1]) == "C") {
				number -= 10
			} else {
				number += 10
			}
		case "L":
			number += 50
		case "C":
			if i < len(s)-1 && (string(s[i+1]) == "D" || string(s[i+1]) == "M") {
				number -= 100
			} else {
				number += 100
			}
		case "D":
			number += 500
		case "M":
			number += 1000
		}
	}
	return number
}
