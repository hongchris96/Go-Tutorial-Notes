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
	fmt.Println("Attempt 1")
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	answerNode := ListNode{}
	answerTracker := &answerNode
	tracker1 := l1
	tracker2 := l2
	carryOver := 0

	for tracker1 != nil && tracker2 != nil {
		currentDigitSum := tracker1.Val + tracker2.Val + carryOver
		if currentDigitSum >= 10 {
			carryOver = 1
		} else {
			carryOver = 0
		}
		answerTracker.Val = currentDigitSum % 10
		tracker1 = tracker1.Next
		tracker2 = tracker2.Next
		if tracker1 == nil && tracker2 == nil {
			break
		}
		nextNode := ListNode{}
		answerTracker.Next = &nextNode
		answerTracker = answerTracker.Next
	}

	var remainingTracker *ListNode

	if tracker1 != nil {
		remainingTracker = tracker1
	} else if tracker2 != nil {
		remainingTracker = tracker2
	}

	for remainingTracker != nil {
		currentDigitSum := remainingTracker.Val + carryOver
		if currentDigitSum >= 10 {
			answerTracker.Val = 0
			if remainingTracker.Next != nil {
				answerTracker.Next = &ListNode{}
				answerTracker = answerTracker.Next
			}
			remainingTracker = remainingTracker.Next
		} else {
			carryOver = 0
			answerTracker.Val = currentDigitSum
			answerTracker.Next = remainingTracker.Next
			break
		}
	}

	if carryOver == 1 {
		answerTracker.Next = &ListNode{Val: 1}
	}

	return &answerNode
}
