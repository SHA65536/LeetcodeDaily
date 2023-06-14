package problem0559

import . "leetcodedaily/helpers/narytree"

/*
Given a n-ary tree, find its maximum depth.
The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
*/

func maxDepth(root *Node) int {
	var res int
	if root == nil {
		return 0
	}
	for i := range root.Children {
		res = max(res, maxDepth(root.Children[i]))
	}
	return res + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
