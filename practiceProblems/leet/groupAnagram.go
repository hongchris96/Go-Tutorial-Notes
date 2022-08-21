package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	groups := map[string][]string{}
	for _, word := range strs {
		letters := strings.Split(word, "")
		sort.Strings(letters)
		sortedWord := strings.Join(letters, "")
		if bucket, ok := groups[sortedWord]; ok {
			groups[sortedWord] = append(bucket, word)
		} else {
			groups[sortedWord] = []string{word}
		}
	}
	answer := make([][]string, len(groups))
	i := 0
	for _, bucket := range groups {
		answer[i] = bucket
		i++
	}
	return answer
}
