package problem0114

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a binary tree, flatten the tree into a "linked list":
The "linked list" should use the same TreeNode class where the right child pointer points to the next node in the list and the left child pointer is always null.
The "linked list" should be in the same order as a pre-order traversal of the binary tree.
*/

func flatten(root *TreeNode) {
	var prev *TreeNode
	for root != nil {
		// If there's no left node, we dont' need to move anything
		if root.Left == nil {
			root = root.Right
			continue
		}
		// Traversing to the end of the left tree, where the right tree should start
		prev = root.Left
		for prev.Right != nil {
			prev = prev.Right
		}
		// Putting the entire right side of the tree and the end of the left tree
		prev.Right = root.Right
		// Moving left side to right side, and nilling the left node
		root.Right = root.Left
		root.Left = nil
		// Moving to next node
		root = root.Right
	}
}
