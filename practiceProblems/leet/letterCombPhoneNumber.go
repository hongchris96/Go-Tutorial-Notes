package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	digitLetter := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
	combs := []string{}
	for _, char := range digits {
		newLetters := digitLetter[string(char)]
		if len(combs) == 0 {
			combs = append(combs, newLetters...)
			continue
		}
		prevCombs := append([]string{}, combs...)
		for i, letter := range newLetters {
			if i == 0 {
				for j := 0; j < len(combs); j++ {
					combs[j] += letter
				}
			} else {
				for j := 0; j < len(prevCombs); j++ {
					combs = append(combs, prevCombs[j]+letter)
				}
			}
		}
	}
	return combs
}
