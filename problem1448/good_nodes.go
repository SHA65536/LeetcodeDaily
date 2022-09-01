package problem1448

import . "leetcodedaily/helpers/binarytree"

/*
Given a binary tree root, a node X in the tree is named good if in the path from root to X there are no nodes with a value greater than X.
Return the number of good nodes in the binary tree.
*/

const minval int = -100000

func goodNodes(root *TreeNode) int {
	return checkNodes(root, minval)
}

func checkNodes(node *TreeNode, max int) int {
	var result int
	if node == nil {
		return 0
	}
	if node.Val >= max {
		result++
		result += checkNodes(node.Left, node.Val)
		result += checkNodes(node.Right, node.Val)
	} else {
		result += checkNodes(node.Left, max)
		result += checkNodes(node.Right, max)
	}
	return result
}
