package problem0112

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a binary tree and an integer targetSum,
return true if the tree has a root-to-leaf path such that adding up all the values along the path equals targetSum.
A leaf is a node with no children.
*/

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return hasPathHelper(root, targetSum)
}

func hasPathHelper(root *TreeNode, target int) bool {
	if root.Left == nil && root.Right == nil {
		return target == root.Val
	}
	if root.Left != nil && hasPathHelper(root.Left, target-root.Val) {
		return true
	}
	if root.Right != nil && hasPathHelper(root.Right, target-root.Val) {
		return true
	}
	return false
}
