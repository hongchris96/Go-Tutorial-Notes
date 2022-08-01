package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println("Merge two sorted list")
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := ListNode{}
	tracker := &head

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tracker.Next = list1
			list1 = list1.Next
		} else {
			tracker.Next = list2
			list2 = list2.Next
		}
		tracker = tracker.Next

	}

	if list1 != nil && list2 == nil {
		tracker.Next = list1
	} else if list1 == nil && list2 != nil {
		tracker.Next = list2
	}

	return head.Next
}
