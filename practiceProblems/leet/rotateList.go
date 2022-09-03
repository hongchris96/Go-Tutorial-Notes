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

	begin := ListNode{}
	end := ListNode{}
	begin.Next = head

	// get length and save pointer to last node
	tracker := head
	length := 0
	for tracker != nil {
		length++
		if tracker.Next == nil {
			end.Next = tracker
		}
		tracker = tracker.Next
	}

	lastNodes := k % length
	if lastNodes == 0 {
		return head
	}

	// find point of pivot
	cutFrom := length - lastNodes - 1
	tracker = head
	for i := 0; i < cutFrom; i++ {
		tracker = tracker.Next
	}
	// performing restructure
	newTail := tracker
	newHead := tracker.Next
	newTail.Next = nil
	end.Next.Next = head
	return newHead

	// ----------------------------------------------------

	// if head == nil {
	// 	return nil
	// }

	// nodes := map[int]*ListNode{}

	// // get length and store nodes
	// tracker := head
	// length := 0
	// for tracker != nil {
	// 	nodes[length] = tracker
	// 	length++
	// 	tracker = tracker.Next
	// }

	// lastNodes := k % length
	// if lastNodes == 0 {
	// 	return head
	// }

	// // performing restructure
	// cutFrom := length - lastNodes - 1
	// answer := nodes[cutFrom+1]
	// nodes[cutFrom].Next = nil
	// nodes[length-1].Next = head

	// return answer
}
