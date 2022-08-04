package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := &ListNode{Next: head}
	answer := prev
	current := head
	next := current.Next
	for next != nil {
		prev.Next = next
		current.Next = next.Next
		next.Next = current

		prev = current
		current = prev.Next
		if current == nil {
			break
		}
		next = current.Next
	}
	return answer.Next
}
