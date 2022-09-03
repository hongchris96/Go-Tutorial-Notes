package main

import "fmt"

func main() {
	fmt.Println()
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	nodes := map[int]*ListNode{}

	// get length and store nodes
	tracker := head
	length := 0
	for tracker != nil {
		nodes[length] = tracker
		length++
		tracker = tracker.Next
	}

	lastNodes := k % length
	if lastNodes == 0 {
		return head
	}

	// performing restructure
	cutFrom := length - lastNodes - 1
	answer := nodes[cutFrom+1]
	nodes[cutFrom].Next = nil
	nodes[length-1].Next = head

	return answer
}
