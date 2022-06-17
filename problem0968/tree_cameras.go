package problem0968

import . "leetcodedaily/helpers/binarytree"

/*
https://leetcode.com/problems/binary-tree-cameras

You are given the root of a binary tree.
We install cameras on the tree nodes where each camera at a node can monitor its parent, itself, and its immediate children.
Return the minimum number of cameras needed to monitor all nodes of the tree.
*/

const (
	NoCamera = iota
	NotNeeded
	HasCamera
)

func minCameraCover(root *TreeNode) int {
	var ans int
	if depthFirst(root, &ans) == NoCamera {
		ans++
	}
	return ans
}

func depthFirst(root *TreeNode, answer *int) int {
	if root == nil {
		return NotNeeded
	}
	left := depthFirst(root.Left, answer)
	right := depthFirst(root.Right, answer)
	if right == NoCamera || left == NoCamera {
		*answer++
		return HasCamera
	}
	if left == HasCamera || right == HasCamera {
		return NotNeeded
	}
	return NoCamera
}
