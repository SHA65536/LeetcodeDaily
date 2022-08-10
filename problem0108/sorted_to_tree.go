package problem0108

import . "leetcodedaily/helpers/binarytree"

/*
Given an integer array nums where the elements are sorted in ascending order, convert it to a height-balanced binary search tree.
A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.
*/

func sortedArrayToBST(nums []int) *TreeNode {
	var mid int = len(nums) / 2
	var head *TreeNode
	// If no numbers, return nil
	if len(nums) == 0 {
		return nil
	}
	// If only 1 number, return it as head
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	head = &TreeNode{Val: nums[mid]}
	// Repeat same on left half
	head.Left = sortedArrayToBST(nums[:mid])
	// Repeat same of right half
	head.Right = sortedArrayToBST(nums[mid+1:])
	return head
}
