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

	cur := &Node{}
	cur.Next = head.Next
	pre := head

	for pre != nil && pre.Next != nil {
		tmp := pre.Next.Next

		pre.Next.Next = pre
		pre.Next = tmp
		pre = tmp

	}

	return cur.Next
}

// 把所有的奇数节点和偶数节点分别排在一起
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd := head
	even := head.Next

	mark := head.Next

	for odd != nil && even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next

		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = mark

	return head
}
