package problem2265

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a binary tree, return the number of nodes where the value of the node is equal to the average of the values in its subtree.
Note:
The average of n elements is the sum of the n elements divided by n and rounded down to the nearest integer.
A subtree of root is a tree consisting of root and all of its descendants.
*/

func averageOfSubtree(root *TreeNode) int {
	var res int
	var dfs func(*TreeNode) (int, int)

	dfs = func(cur *TreeNode) (int, int) {
		if cur == nil {
			return 0, 0
		}
		lS, lC := dfs(cur.Left)
		rS, rC := dfs(cur.Right)
		if cur.Val == (lS+rS+cur.Val)/(lC+rC+1) {
			res++
		}
		return (lS + rS + cur.Val), (lC + rC + 1)
	}

	dfs(root)

	return res
}
