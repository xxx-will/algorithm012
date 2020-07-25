/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre := head
	var cur *ListNode

	for pre != nil {

		tmp := pre.Next

		pre.Next = cur
		cur = pre

		pre = tmp
	}

	return cur
}

// @lc code=end

