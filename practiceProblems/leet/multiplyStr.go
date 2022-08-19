package main

import "fmt"

func main() {
	fmt.Println(multiply("498828660196", "840477629533")) // "419254329864656431168468"
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	answer := []int{}
	// performing math operation one digit at a time right to left
	carryOver := 0
	for i := len(num1) - 1; i > -1; i-- {
		rune1 := num1[i]
		carryOver = 0
		for j := len(num2) - 1; j > -1; j-- {
			rune2 := num2[j]
			// getting product of current two digits
			digProd := int(rune2-48)*int(rune1-48) + carryOver
			// adding product to the right place in the answer
			answerDigPosition := len(num1) + len(num2) - j - i - 2
			// either adding to existing slot or append to a new slot
			if answerDigPosition < len(answer) {
				digProd += answer[answerDigPosition]
				// slot only holds single digit integers
				// update carryover to next operation
				if digProd > 9 {
					answer[answerDigPosition] = digProd % 10
					carryOver = digProd / 10
				} else {
					answer[answerDigPosition] = digProd
					carryOver = 0
				}
			} else {
				if digProd > 9 {
					answer = append(answer, digProd%10)
					carryOver = digProd / 10
				} else {
					answer = append(answer, digProd)
					carryOver = 0
				}
			}
		}
		// append a new slot for carryover after each inner loop
		if carryOver > 0 {
			answer = append(answer, carryOver)
		}
	}
	// convert slice of int to string
	stringAnswer := ""
	for _, num := range answer {
		stringAnswer = string(num+48) + stringAnswer
	}
	return stringAnswer

}
