package problem0951

import . "leetcodedaily/helpers/binarytree"

/*
For a binary tree T, we can define a flip operation as follows: choose any node, and swap the left and right child subtrees.
A binary tree X is flip equivalent to a binary tree Y if and only if we can make X equal to Y after some number of flip operations.
Given the roots of two binary trees root1 and root2, return true if the two trees are flip equivalent or false otherwise.
*/

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if (root1 == nil) != (root2 == nil) {
		return false
	}
	if root1 == nil {
		return true
	}
	if root1.Val != root2.Val {
		return false
	}
	if flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right) {
		return true
	}
	if flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left) {
		return true
	}
	return false
}
