package problem0094

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a binary tree, return the inorder traversal of its nodes' values.
*/

func inorderTraversal(root *TreeNode) []int {
	var results = []int{}
	traverse(root, &results)
	return results
}

func traverse(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		traverse(root.Left, res)
	}
	*res = append(*res, root.Val)
	if root.Right != nil {
		traverse(root.Right, res)
	}
}
