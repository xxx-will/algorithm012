package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	nMap := make(map[int]int)
	for i, n := range inorder {
		nMap[n] = i
	}

	return build(preorder, 0, len(preorder)-1, nMap, 0, len(inorder)-1)

}

func build(preorder []int, pLeft, pRight int, nMap map[int]int, iLeft, iRight int) *TreeNode {
	if pLeft > pRight {
		return nil
	}

	root := &TreeNode{
		Val: preorder[pLeft],
	}

	rootIndex := nMap[root.Val]

	root.Left = build(preorder, pLeft+1, rootIndex-iLeft+pLeft, nMap, iLeft, rootIndex-1)

	root.Right = build(preorder, rootIndex-iLeft+pLeft+1, pRight, nMap, rootIndex+1, iRight)

	return root
}
