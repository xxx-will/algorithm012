package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//mergeTwoLists  链表合并
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	head := &ListNode{}
	cur := head

	for l1.Next != nil && l2.Next != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1.Next == nil {
		cur.Next = l2
	}
	if l2.Next == nil {
		cur.Next = l1
	}

	return head.Next
}
