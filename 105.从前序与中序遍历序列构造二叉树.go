/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) != len(inorder) {
		return nil
	}

	indexMap := make(map[int]int)

	for i, v := range inorder {
		indexMap[v] = i
	}

	return build(indexMap, preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func build(indexMap map[int]int, pre []int, pre_start, pre_end int, in []int, in_start, in_end int) *TreeNode {
	if pre_start > pre_end {
		return nil
	}

	root := &TreeNode{}
	root.Val = pre[pre_start]

	in_index := indexMap[root.Val]
	size := in_index - 1 - in_start

	root.Left = build(indexMap, pre, pre_start+1, pre_start+1+size, in, in_start, in_index-1)
	root.Right = build(indexMap, pre, pre_start+1+size+1, pre_end, in, in_index+1, in_end)

	return root
}

// @lc code=end

