/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) != len(inorder) {
		return nil
	}

	indexMap := make(map[int]int)

	for i, v := range inorder {
		indexMap[v] = i
	}

	return build(indexMap, inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func build(indexMap map[int]int, inorder []int, iStart, iEnd int, postorder []int, pStart, pEnd int) *TreeNode {
	if pStart > pEnd {
		return nil
	}

	root := &TreeNode{
		Val: postorder[pEnd],
	}

	inIndex := indexMap[root.Val]
	size := inIndex - 1 - iStart

	root.Left = build(indexMap, inorder, iStart, inIndex-1, postorder, pStart, pStart+size)
	root.Right = build(indexMap, inorder, inIndex+1, iEnd, postorder, pStart+size+1, pEnd-1)

	return root
}

// @lc code=end

