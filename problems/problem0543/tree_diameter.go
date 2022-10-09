package problem0543

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a binary tree, return the length of the diameter of the tree.
The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.
The length of a path between two nodes is represented by the number of edges between them.
*/

func diameterOfBinaryTree(root *TreeNode) int {
	var res int
	findDiameter(root, &res)
	return res
}

func findDiameter(root *TreeNode, res *int) int {
	if root == nil {
		return -1
	}
	left := findDiameter(root.Left, res) + 1
	right := findDiameter(root.Right, res) + 1
	if left+right > *res {
		*res = left + right
	}
	if left > right {
		return left
	}
	return right
}
