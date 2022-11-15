package problem0222

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a complete binary tree, return the number of the nodes in the tree.
According to Wikipedia, every level, except possibly the last, is completely filled in a complete binary tree,
and all nodes in the last level are as far left as possible.
It can have between 1 and 2h nodes inclusive at the last level h.
Design an algorithm that runs in less than O(n) time complexity.
*/

func countNodes(root *TreeNode) int {
	var left, right int
	var temp *TreeNode
	if root == nil {
		return 0
	}
	for temp = root; temp != nil; temp = temp.Left {
		left++
	}
	for temp = root; temp != nil; temp = temp.Right {
		right++
	}
	if left == right {
		return (1 << left) - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

func countNodesNaive(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}
