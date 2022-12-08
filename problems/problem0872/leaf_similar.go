package problem0872

import . "leetcodedaily/helpers/binarytree"

/*
Consider all the leaves of a binary tree, from left to right order, the values of those leaves form a leaf value sequence.
*/

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var leaf1 = []int{}
	var leaf2 = []int{}
	getSequence(root1, &leaf1)
	getSequence(root2, &leaf2)
	if len(leaf1) != len(leaf2) {
		return false
	}
	for i := range leaf1 {
		if leaf1[i] != leaf2[i] {
			return false
		}
	}
	return true
}

func getSequence(root *TreeNode, res *[]int) {
	if root.Left == nil && root.Right == nil {
		(*res) = append((*res), root.Val)
	}
	if root.Left != nil {
		getSequence(root.Left, res)
	}
	if root.Right != nil {
		getSequence(root.Right, res)
	}
}
