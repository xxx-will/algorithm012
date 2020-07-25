/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	//边界情况
	// 1=m n=长度
	// m=n

	if head.Next == nil {
		return head
	}

	i := 1
	pre := head

	var start, end, cur *ListNode

	for i < m && pre != nil {
		cur = pre
		pre = pre.Next
		i++
	}
	start = cur
	end = pre
	cur = nil

	for i <= n && pre != nil {
		tmp := pre.Next
		pre.Next = cur
		cur = pre
		pre = tmp
		i++
	}

	start.Next = cur
	end.Next = pre

	return head
}

// @lc code=end

