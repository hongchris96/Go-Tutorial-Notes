package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("aaaa"))
	fmt.Println(longestPalindrome("caad"))
	fmt.Println(longestPalindrome("babad"))
}

func longestPalindrome(s string) string {

	globalMax := string(s[0])
	i := 0

	for ; i < len(s)-1; i++ {
		leftIdx := i
		rightIdx := i
		if i-1 > -1 && s[i-1] == s[i+1] {
			leftIdx = i - 1
			rightIdx = i + 1
			for leftIdx > -1 && rightIdx < len(s) && s[leftIdx] == s[rightIdx] {
				leftIdx -= 1
				rightIdx += 1
			}
			if rightIdx-leftIdx-1 > len(globalMax) {
				globalMax = string(s[leftIdx+1 : rightIdx])
			}
		}
		if s[i] == s[i+1] {
			leftIdx = i
			rightIdx = i + 1
			for leftIdx > -1 && rightIdx < len(s) && s[leftIdx] == s[rightIdx] {
				leftIdx -= 1
				rightIdx += 1
			}
			if rightIdx-leftIdx-1 > len(globalMax) {
				globalMax = string(s[leftIdx+1 : rightIdx])
			}
		}
	}
	return globalMax
}
