package main

type ListNode struct {
	val  int
	next *ListNode
}

func printListFromTailToHead(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	ph := head
	var p *ListNode
	for ph != nil {
		if p == nil {
			p = &ListNode{ph.val, nil}
		} else {
			p1 := &ListNode{ph.val, nil}
			p1.next = p
			p = p1
		}
		ph = ph.next
	}

	return p

}
