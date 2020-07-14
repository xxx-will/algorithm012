package main

type Node struct {
	Value int
	Next  *Node
}

// 单向链表反转
func reverseOneway(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	l := &Node{}
	cur := head
	l.Next = nil

	for cur != nil {

		tmp := cur.Next
		cur.Next = l.Next
		l.Next = cur
		cur = tmp
	}

	return l.Next
}

func reverseList(head *Node) *Node {

	var cur *Node
	pre := head

	for pre != nil {
		t := pre.Next
		pre.Next = cur
		cur = pre
		pre = t
	}
	return cur

}

// 奇偶位交换

func exchangeOddEven(head *Node) *Node {

}
