package problem1026

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a binary tree, find the maximum value v for which there exist different nodes a and b where v = |a.val - b.val| and a is an ancestor of b.
A node a is an ancestor of b if either: any child of a is equal to b or any child of a is an ancestor of b.
*/

func maxAncestorDiff(root *TreeNode) int {
	var res int
	maxAncestor(&res, root.Val, root.Val, root)
	return res
}

func maxAncestor(res *int, maxV, minV int, root *TreeNode) {
	maxV = max(root.Val, maxV)
	minV = min(root.Val, minV)
	*res = max(*res, diff(maxV, minV))
	if root.Left != nil {
		maxAncestor(res, maxV, minV, root.Left)
	}
	if root.Right != nil {
		maxAncestor(res, maxV, minV, root.Right)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func diff(a, b int) int {
	a = a - b
	if a < 0 {
		return -a
	}
	return a
}
