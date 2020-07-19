package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归

func inorderTraversal(root *TreeNode) []int {

	arr1 := make([]int, 0)

	// 中序遍历
	traversal1(root, &arr1)

	return arr1
}

func traversal1(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	if root.Left != nil {
		traversal1(root.Left, arr)
	}

	*arr = append(*arr, root.Val)

	if root.Right != nil {
		traversal1(root.Right, arr)
	}

}

//
func preorderTraversal(root *TreeNode) []int {

	arr2 := make([]int, 0)

	// 前序遍历
	traversal2(root, &arr2)

	return arr2
}

func traversal2(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	*arr = append(*arr, root.Val)

	if root.Left != nil {
		traversal2(root.Left, arr)
	}
	if root.Right != nil {
		traversal2(root.Right, arr)
	}
	return
}
