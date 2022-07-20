package main

import "fmt"

func main() {
	s := "bbbbb"
	s2 := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring((s)))
	fmt.Println(lengthOfLongestSubstring((s2)))

}

func lengthOfLongestSubstring(s string) int {
	// maxLength := 0
	// uniqueLetters := map[string]int{}

	// for i := 0; i < len(s); i++ {
	// 	if idx, ok := uniqueLetters[string(s[i])]; ok {
	// 		for k, v := range uniqueLetters {
	// 			if v < idx {
	// 				delete(uniqueLetters, k)
	// 			}
	// 		}
	// 	}
	// 	uniqueLetters[string(s[i])] = i
	// 	currentMapLength := len(uniqueLetters)
	// 	if currentMapLength > maxLength {
	// 		maxLength = currentMapLength
	// 	}
	// }
	// return maxLength

	// --------------------------------------------------------------------------

	// maxLength := 0
	// charLastPosition := map[string]int{}
	// iHead := 0
	// iTail := 0
	// for iTail < len(s) {
	//     if lastPosition, exist := charLastPosition[string(s[iTail])]; exist {
	//         if lastPosition + 1 > iHead {
	//             iHead = lastPosition + 1
	//         }
	//     }
	//     charLastPosition[string(s[iTail])] = iTail
	//     if currentLength := iTail + 1 - iHead; currentLength > maxLength {
	//         maxLength = currentLength
	//     }
	//     iTail++
	// }
	// return maxLength

	// --------------------------------------------------------------------------

	if len(s) == 0 {
		return 0
	}
	maxLength := 0
	iHead := 0
	iTail := 0
	charLastPosition := make(map[byte]int)
	for iTail < len(s) {
		if lastPos, exist := charLastPosition[s[iTail]]; exist {
			if currentLength := iTail - iHead; currentLength > maxLength {
				maxLength = currentLength
			}
			// delete all the letter leading up to last Position
			for i := iHead; i <= lastPos; i++ {
				delete(charLastPosition, s[i])
			}
			iHead = lastPos + 1
		}

		charLastPosition[s[iTail]] = iTail
		iTail++
	}
	if currentLength := iTail - iHead; currentLength > maxLength {
		maxLength = currentLength
	}
	return maxLength
}
