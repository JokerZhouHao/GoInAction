package main

type ListNode struct {
	val  int
	next *ListNode
}

func FindKthToTail(head *ListNode, k int) *ListNode {
	p1, p2 := head, head
	for p1 != nil && k > 0 {
		p1 = p1.next
		k--
	}
	if p1 == nil && k != 0 {
		return nil
	}

	for p1 != nil {
		p1 = p1.next
		p2 = p2.next
	}
	return p2
}
