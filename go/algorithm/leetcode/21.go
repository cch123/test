package main

func main() {
}

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head = &ListNode{}
	var currentNode = head
	for l1 != nil || l2 != nil {
		if l1 == nil {
			for l2 != nil {
				tmpNode := &ListNode{Val: l2.Val}
				l2 = l2.Next
				currentNode.Next = tmpNode
				currentNode = currentNode.Next
			}
			break
		}

		if l2 == nil {
			for l1 != nil {
				tmpNode := &ListNode{Val: l1.Val}
				l1 = l1.Next
				currentNode.Next = tmpNode
				currentNode = currentNode.Next
			}
			break
		}

		if l1.Val < l2.Val {
			tmpNode := &ListNode{Val: l1.Val}
			l1 = l1.Next
			currentNode.Next = tmpNode
			currentNode = currentNode.Next
		} else {
			tmpNode := &ListNode{Val: l2.Val}
			l2 = l2.Next
			currentNode.Next = tmpNode
			currentNode = currentNode.Next
		}

	}
	return head.Next
}
