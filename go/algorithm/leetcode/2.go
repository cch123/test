package main

import "fmt"

func main() {
	l1 := &ListNode{Val: 5}
	l2 := &ListNode{Val: 5}
	l := addTwoNumbers(l1, l2)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	reserved := 0
	var listHead = &ListNode{}
	var currentNode = listHead
	for l1 != nil {
		if l2 != nil {
			currentNode.Next = &ListNode{}
			currentNode = currentNode.Next

			currentNode.Val = (l1.Val + l2.Val + reserved) % 10
			reserved = (l1.Val + l2.Val + reserved) / 10
			l1 = l1.Next
			l2 = l2.Next
		} else {
			// l1!=nil && l2=nil
			currentNode.Next = &ListNode{}
			currentNode = currentNode.Next

			currentNode.Val = (l1.Val + reserved) % 10
			reserved = (l1.Val + reserved) / 10
			l1 = l1.Next
		}
	}

	for l2 != nil {
		currentNode.Next = &ListNode{}
		currentNode = currentNode.Next

		currentNode.Val = (l2.Val + reserved) % 10
		reserved = (l2.Val + reserved) / 10
		l2 = l2.Next
	}

	if reserved > 0 {
		currentNode.Next = &ListNode{}
		currentNode = currentNode.Next
		currentNode.Val = reserved
	}

	return listHead.Next
}
