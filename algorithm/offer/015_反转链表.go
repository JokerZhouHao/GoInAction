package main

type ListNode struct {
	val  int
	next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	p1, p2 := (*ListNode)(nil), head
	for p2 != nil {
		p3 := p2.next
		p2.next = p1
		p1 = p2
		p2 = p3
	}
	return p1
}
