package main

import "fmt"

func main() {
	fmt.Println(strStr("hello", "ll"))
}

func strStr(haystack string, needle string) int {
	needleLen := len(needle)
	if needleLen == 0 {
		return 0
	}
	if needleLen > len(haystack) {
		return -1
	}
	for i, charCode := range haystack[:len(haystack)-needleLen+1] {
		if byte(charCode) == needle[0] && haystack[i:i+needleLen] == needle {
			return i
		}
	}
	return -1
}
