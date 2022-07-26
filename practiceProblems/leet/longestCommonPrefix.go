package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{"ab", "a"}))
}

func longestCommonPrefix(strs []string) string {
	commonPrefix := strs[0]
	if len(strs[0]) == 0 {
		return ""
	}

	for i, word := range strs {
		if i == 0 {
			continue
		}
		for j := 0; j < len(commonPrefix); j++ {
			if j == len(word) || word[j] != commonPrefix[j] {
				commonPrefix = word[:j]
				break
			}
		}
	}
	return commonPrefix
}
