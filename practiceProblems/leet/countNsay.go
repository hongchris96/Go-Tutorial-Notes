package main

import "fmt"

func main() {
	fmt.Println(countAndSay(4))
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	prevResults := countAndSay(n - 1)
	finalSay := ""
	var currentNum rune
	count := 0
	for i, numRune := range prevResults {
		if i == 0 {
			currentNum = numRune
			count++
			if i == len(prevResults)-1 {
				finalSay += string(count+48) + string(currentNum)
			}
			continue
		}
		if numRune == currentNum {
			count++
		} else {
			finalSay += string(count+48) + string(currentNum)
			currentNum = numRune
			count = 1
		}
		if i == len(prevResults)-1 {
			finalSay += string(count+48) + string(currentNum)
		}
	}
	return finalSay
}
