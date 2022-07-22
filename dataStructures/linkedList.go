package main

import "fmt"

type node struct {
	next *node
	data int
}

type linkedList struct {
	head   *node
	length int
}

func main() {
	fmt.Println("Linked List:")
	myList := linkedList{}
	node1 := &node{data: 48} // needs to pass in a pointer here
	node2 := &node{data: 47}
	node3 := &node{data: 46}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	fmt.Println(myList)
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}
