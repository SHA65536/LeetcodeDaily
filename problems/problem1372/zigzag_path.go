package problem1372

import . "leetcodedaily/helpers/binarytree"

/*
You are given the root of a binary tree.
A ZigZag path for a binary tree is defined as follow:
    Choose any node in the binary tree and a direction (right or left).
    If the current direction is right, move to the right child of the current node; otherwise, move to the left child.
    Change the direction from right to left or from left to right.
    Repeat the second and third steps until you can't move in the tree.
Zigzag length is defined as the number of nodes visited - 1. (A single node has a length of 0).
Return the longest ZigZag path contained in that tree.
*/

type direction int

const (
	LEFT direction = iota
	RIGHT
)

func longestZigZag(root *TreeNode) int {
	var res int
	dfs(root.Left, 0, LEFT, &res)
	dfs(root.Right, 0, RIGHT, &res)
	return res
}

func dfs(root *TreeNode, depth int, d direction, res *int) {
	*res = max(*res, depth)
	if root == nil {
		return
	}

	if d == RIGHT {
		dfs(root.Left, depth+1, LEFT, res)
		dfs(root.Right, 0, RIGHT, res)
	} else {
		dfs(root.Left, 0, LEFT, res)
		dfs(root.Right, depth+1, RIGHT, res)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
