package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println("Attempt 1")
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	emptyHead := ListNode{Next: head}
	prevTracker := &emptyHead

	for prevTracker != nil {
		currentTracker := prevTracker.Next
		potentialNewNext := currentTracker.Next
		for i := 0; i < n-1; i++ {
			currentTracker = currentTracker.Next
		}
		if currentTracker.Next == nil {
			prevTracker.Next = potentialNewNext
			break
		}
		prevTracker = prevTracker.Next
	}

	return emptyHead.Next
}
