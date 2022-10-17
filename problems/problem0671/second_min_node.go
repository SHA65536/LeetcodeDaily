package problem0671

import . "leetcodedaily/helpers/binarytree"

/*
Given a non-empty special binary tree consisting of nodes with the non-negative value,
where each node in this tree has exactly two or zero sub-node.
If the node has two sub-nodes, then this node's value is the smaller value among its two sub-nodes.
More formally, the property root.val = min(root.left.val, root.right.val) always holds.
Given such a binary tree, you need to output the second minimum value in the set made of all the nodes' value in the whole tree.
If no such second minimum value exists, output -1 instead.
*/

func findSecondMinimumValue(root *TreeNode) int {
	var left, right int
	// If no children, don't bother checking
	if root.Left == nil {
		return -1
	}

	// find first value that is bigger than current value
	left = root.Left.Val
	if root.Left.Val == root.Val {
		left = findSecondMinimumValue(root.Left)
	}

	// find first value that is bigger than current value
	right = root.Right.Val
	if root.Right.Val == root.Val {
		right = findSecondMinimumValue(root.Right)
	}

	if left != -1 && right != -1 {
		// If both are different from current, take the smaller
		return min(left, right)
	} else {
		// If one of them is the same, take the other one
		return max(left, right)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
