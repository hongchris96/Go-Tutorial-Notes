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
	answerTracker := answerNode
	tracker1 := l1
	tracker2 := l2
	carryOver := 0
	// fmt.Println(answerTracker.Val)
	// fmt.Println((tracker1).Val, (*tracker2).Val)
	// tracker1 = (tracker1).Next
	// fmt.Println((tracker1).Val, (*tracker2).Val)
	// tracker1 = tracker1.Next
	// fmt.Println(tracker1 == nil)
	// tracker1 = tracker1.Next
	// fmt.Println(tracker1 != nil, (*tracker2).Val)

	for tracker1 != nil || tracker2 != nil {
		currentDigitSum := tracker1.Val + tracker2.Val
		if currentDigitSum > 10 {
			carryOver = 1
		}
		answerNode.Val = currentDigitSum % 10
		tracker1 = tracker1.Next
		tracker2 = tracker2.Next
		nextNode := ListNode{}
		(*answerTracker).Next = &nextNode
		answerTracker = answerTracker.Next
	}

	return &answerNode
}
