package main

import "fmt"

func main() {
	fmt.Println(multiply("498828660196", "840477629533")) // "419254329864656431168468"
}

func multiply(num1 string, num2 string) string {
	int1 := 0
	int2 := 0
	for i, dig := range num1 {
		int1 += tenPow(len(num1)-1-i) * int(dig-48)
	}
	for i, dig := range num2 {
		int2 += tenPow(len(num2)-1-i) * int(dig-48)
	}
	prod := int1 * int2
	fmt.Println(prod)
	if prod == 0 {
		return "0"
	}
	answer := ""
	for prod > 0 {
		dig := prod % 10
		answer = string(dig+48) + answer
		prod /= 10
	}
	return answer
}

func tenPow(pow int) int {
	answer := 1
	for i := 0; i < pow; i++ {
		answer *= 10
	}
	return answer
}
