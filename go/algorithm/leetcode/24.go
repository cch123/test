package main

func main() {
	l := &ListNode{Val: 1}
	l.Next = &ListNode{Val: 2}
	l.Next.Next = &ListNode{Val: 3}
	l.Next.Next.Next = &ListNode{Val: 4}
	h := l
	for h != nil {
		h = h.Next
	}
	l = swapPairs(l)
	for l != nil {
		l = l.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	leftNode := head
	rightNode := head.Next
	head = rightNode
	for {
		leftNode.Next, rightNode.Next = rightNode.Next, leftNode
		leftNode, rightNode = rightNode, leftNode

		//如果right.right!=nil，要先把prev的next改成right.right
		prev := rightNode
		leftNode = rightNode.Next
		if leftNode == nil {
			break
		}
		rightNode = leftNode.Next
		if rightNode == nil {
			break
		}
		prev.Next = rightNode
	}

	return head
}
