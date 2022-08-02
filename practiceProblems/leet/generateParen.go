package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3)) // "((()))","(()())","(())()","()(())","()()()"
	fmt.Println(generateParenthesis(1)) // "()"
}

func generateParenthesis(n int) []string {
	return parenthesisBuilder("", n, n)
}

func parenthesisBuilder(partial string, opening, closing int) []string {
	if opening == 0 && closing == 1 {
		return []string{partial + ")"}
	}
	answer := []string{}
	if opening > 0 {
		answer = append(answer, parenthesisBuilder(partial+"(", opening-1, closing)...)
	}
	if opening < closing {
		answer = append(answer, parenthesisBuilder(partial+")", opening, closing-1)...)
	}
	return answer
}
