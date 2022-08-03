package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println("Merge K Sorted List")
}

func mergeKLists(lists []*ListNode) *ListNode {
	head := ListNode{}
	tracker := &head
	allNil := false
	for !allNil {
		allNil = true
		currentMinValue := getMinValFromHeads(lists)
		for i, node := range lists {
			if node != nil {
				allNil = false
				if node.Val == currentMinValue {
					tracker.Next = node
					tracker = tracker.Next
					lists[i] = node.Next
				}
			}
		}
	}
	return head.Next
}

func getMinValFromHeads(lists []*ListNode) int {
	min := math.MaxInt64
	for _, node := range lists {
		if node != nil && node.Val < min {
			min = node.Val
		}
	}
	return min
}
