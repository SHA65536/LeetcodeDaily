package problem0572

import . "leetcodedaily/helpers/binarytree"

/*
Given the roots of two binary trees root and subRoot,
return true if there is a subtree of root with the same structure and node values of subRoot and false otherwise.
A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants.
The tree tree could also be considered as a subtree of itself.
*/

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return checkHere(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

// Checks if two trees match
func checkHere(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if (root != nil && subRoot == nil) || (root == nil && subRoot != nil) {
		return false
	}
	if root.Val != subRoot.Val {
		return false
	}
	return checkHere(root.Left, subRoot.Left) && checkHere(root.Right, subRoot.Right)
}
