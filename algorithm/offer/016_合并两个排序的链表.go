package main

type ListNode struct {
	val  int
	next *ListNode
}

func Merge(list1, list2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	p := head
	for list1 != nil && list2 != nil {
		if list1.val <= list2.val {
			p.next = list1
			list1 = list1.next
		} else {
			p.next = list2
			list2 = list2.next
		}
	}
	if list1 == nil {
		p.next = list2
	}
	if list2 == nil {
		p.next = list1
	}
	return head.next
}
