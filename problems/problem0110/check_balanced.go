package problem0110

import . "leetcodedaily/helpers/binarytree"

/*
Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as:
A binary tree in which the left and right subtrees of every node differ in height by no more than 1.
*/

func isBalanced(root *TreeNode) bool {
	_, ok := treeHeight(root)
	return ok
}

func treeHeight(root *TreeNode) (int, bool) {
	var left, right int
	var ok bool
	if root == nil {
		return 0, true
	}
	if left, ok = treeHeight(root.Left); !ok {
		return 0, false
	}
	if right, ok = treeHeight(root.Right); !ok {
		return 0, false
	}
	if diff(left, right) > 1 {
		return 0, false
	}
	return max(left, right) + 1, true
}

func max(a, b int) int {
	if a > b {
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
