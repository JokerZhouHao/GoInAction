package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, p, p1 *ListNode
	k := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + k
		p1 = &ListNode{sum % 10, nil}
		k = sum / 10
		if nil == head {
			head = p1
			p = p1
		} else {
			p.Next = p1
			p = p.Next
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 == nil {
		p1 = l2
	} else {
		p1 = l1
	}
	for p1 != nil {
		sum := p1.Val + k
		p.Next = &ListNode{sum % 10, nil}
		k = sum / 10
		p = p.Next
		p1 = p1.Next
	}
	if k != 0 {
		p.Next = &ListNode{k, nil}
	}
	return head
}
