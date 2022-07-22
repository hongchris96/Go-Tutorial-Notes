/**
 * Definition for singly-linked list.
 */
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println("Attempt 2")
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// answerNode := ListNode{}
	// answerTracker := &answerNode
	// tracker1 := l1
	// tracker2 := l2
	// carryOver := 0

	// for tracker1 != nil && tracker2 != nil {
	// 	currentDigitSum := tracker1.Val + tracker2.Val + carryOver
	// 	if currentDigitSum >= 10 {
	// 		carryOver = 1
	// 	} else {
	// 		carryOver = 0
	// 	}
	// 	answerTracker.Val = currentDigitSum % 10
	// 	tracker1 = tracker1.Next
	// 	tracker2 = tracker2.Next
	// 	if tracker1 == nil && tracker2 == nil {
	// 		break
	// 	}
	// 	nextNode := ListNode{}
	// 	answerTracker.Next = &nextNode
	// 	answerTracker = answerTracker.Next
	// }

	// var remainingTracker *ListNode

	// if tracker1 != nil {
	// 	remainingTracker = tracker1
	// } else if tracker2 != nil {
	// 	remainingTracker = tracker2
	// }

	// for remainingTracker != nil {
	// 	currentDigitSum := remainingTracker.Val + carryOver
	// 	if currentDigitSum >= 10 {
	// 		answerTracker.Val = 0
	// 		if remainingTracker.Next != nil {
	// 			answerTracker.Next = &ListNode{}
	// 			answerTracker = answerTracker.Next
	// 		}
	// 		remainingTracker = remainingTracker.Next
	// 	} else {
	// 		carryOver = 0
	// 		answerTracker.Val = currentDigitSum
	// 		answerTracker.Next = remainingTracker.Next
	// 		break
	// 	}
	// }

	// if carryOver == 1 {
	// 	answerTracker.Next = &ListNode{Val: 1}
	// }

	// return &answerNode

	// ----------------------------------------------------------------------

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	headNode := ListNode{}
	answerTracker := &headNode
	carryOver := 0

	// do everything in one loop
	for !(l1 == nil && l2 == nil && carryOver == 0) {
		num1 := 0
		num2 := 0

		if l1 != nil {
			num1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num2 = l2.Val
			l2 = l2.Next
		}

		currentDigitSum := num1 + num2 + carryOver
		if currentDigitSum > 9 {
			currentDigitSum -= 10
			carryOver = 1
		} else {
			carryOver = 0
		}

		answerTracker.Next = &ListNode{Val: currentDigitSum}
		answerTracker = answerTracker.Next
	}

	return headNode.Next
}
