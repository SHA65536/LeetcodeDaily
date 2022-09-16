package problem0098

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a binary tree, determine if it is a valid binary search tree (BST).
A valid BST is defined as follows:
The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.
*/

func isValidBST(root *TreeNode) bool {
	var searchOrder = []int{}
	traverseTree(root, &searchOrder)
	for i := 1; i < len(searchOrder); i++ {
		if searchOrder[i-1] >= searchOrder[i] {
			return false
		}
	}
	return true
}

func traverseTree(root *TreeNode, list *[]int) {
	if root.Left != nil {
		traverseTree(root.Left, list)
	}
	*list = append(*list, root.Val)
	if root.Right != nil {
		traverseTree(root.Right, list)
	}
}
