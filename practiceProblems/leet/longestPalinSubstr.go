package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("aaaa"))
	fmt.Println(longestPalindrome("caad"))
	fmt.Println(longestPalindrome("babad"))
}

func longestPalindrome(s string) string {

	globalMax := string(s[0])

	for i := 0; i < len(s)-1; i++ {
		// odd number palindrome
		if i-1 > -1 && s[i-1] == s[i+1] {
			middleOutUpdateLongestPalindrome(i-1, i+1, s, &globalMax)
		}
		// even number palindrome
		if s[i] == s[i+1] {
			middleOutUpdateLongestPalindrome(i, i+1, s, &globalMax)
		}
	}
	return globalMax
}

func middleOutUpdateLongestPalindrome(left, right int, str string, globalLongest *string) {
	// expand left and right until out of range or no longer a palindrome
	for left > -1 && right < len(str) && str[left] == str[right] {
		left -= 1
		right += 1
	}
	// update globalMax
	if right-left-1 > len(*globalLongest) {
		*globalLongest = string(str[left+1 : right])
	}
}
