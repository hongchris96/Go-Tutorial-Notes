package main

import "fmt"

func main() {
	fmt.Println(isValid("(([{}]))"))
	fmt.Println(isValid("(){}[]"))
	fmt.Println(isValid("(()({}])"))
}

func isValid(s string) bool {
	opening := map[rune]rune{
		40:  41,
		123: 125,
		91:  93,
	}
	closing := map[rune]rune{
		41:  40,
		125: 123,
		93:  91,
	}
	stack := []rune{}
	for _, charCode := range s {
		if _, ok := opening[charCode]; ok {
			stack = append(stack, charCode)
		} else if len(stack) > 0 && stack[len(stack)-1] == closing[charCode] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
