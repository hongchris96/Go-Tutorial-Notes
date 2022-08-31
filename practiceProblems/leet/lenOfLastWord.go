package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("   fly me   to   the moon   ")) // 4
	fmt.Println(lengthOfLastWord("Hello World"))                  // 5
}

func lengthOfLastWord(s string) int {
	count := 0
	for i := len(s) - 1; i > -1; i-- {
		if string(s[i]) != " " {
			count += 1
		} else if count > 0 {
			return count
		}
	}
	return count
}
