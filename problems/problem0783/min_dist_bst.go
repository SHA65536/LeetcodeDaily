package problem0783

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a Binary Search Tree (BST),
return the minimum difference between the values of any two different nodes in the tree.
*/

func minDiffInBST(root *TreeNode) int {
	var prev, res int = -10001, 10001
	traverse(root, &prev, &res)
	return res
}

func traverse(root *TreeNode, prev, res *int) {
	if root == nil {
		return
	}
	traverse(root.Left, prev, res)
	*res = min(*res, root.Val-*prev)
	*prev = root.Val
	traverse(root.Right, prev, res)
}

func minDiffInBSTStack(root *TreeNode) int {
	var res int = 10001
	var prev *TreeNode
	var stack = []*TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// Pop
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil {
			res = min(res, root.Val-prev.Val)
		}
		prev = root
		root = root.Right
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
