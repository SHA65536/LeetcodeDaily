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
	Covered
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
	// Leaves will need to be covered so we mark
	// their children as covered
	if root == nil {
		return Covered
	}
	left := depthFirst(root.Left, answer)
	right := depthFirst(root.Right, answer)
	// If the childeren need to be covered, we add a camera
	// to the current leaf
	if right == NoCamera || left == NoCamera {
		*answer++
		return HasCamera
	}
	// If one of the childeren has a camera
	// we mark current leaf as covered
	if left == HasCamera || right == HasCamera {
		return Covered
	}
	return NoCamera
}
