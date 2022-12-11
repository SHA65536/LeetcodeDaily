package problem0124

import (
	. "leetcodedaily/helpers/binarytree"
	"math"
)

/*
A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them.
A node can only appear in the sequence at most once. Note that the path does not need to pass through the root.
The path sum of a path is the sum of the node's values in the path.
Given the root of a binary tree, return the maximum path sum of any non-empty path.
*/

func maxPathSum(root *TreeNode) int {
	var maxSum int = math.MinInt
	MaxPathHelper(root, &maxSum)
	return maxSum
}

func MaxPathHelper(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}
	left := MaxPathHelper(root.Left, maxSum)
	right := MaxPathHelper(root.Right, maxSum)
	// max path without going up
	maxSubPath := max(root.Val+max(left, right), root.Val)
	*maxSum = max3(*maxSum, maxSubPath, root.Val+left+right)
	return maxSubPath
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func max3(a, b, c int) int {
	return max(max(a, b), c)
}
