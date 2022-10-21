package problem0226

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a binary tree, invert the tree, and return its root.
*/

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
