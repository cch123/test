package main

import "fmt"

func main() {
	var a = &ListNode{}
	a.Val = 1
	var b = &ListNode{Val: 2}
	a.Next = b
	fmt.Println(removeNthFromEnd(a, 1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 1 && head.Next == nil {
		return nil
	}
	// double pointer
	var startPoint = head
	var end = head
	for i := 0; i < n; i++ {
		end = end.Next
		if end == nil {
			return head.Next
		}
	}

	for end.Next != nil {
		end = end.Next
		startPoint = startPoint.Next
	}

	//delete startPoint's next
	startPoint.Next = startPoint.Next.Next
	return head

}
