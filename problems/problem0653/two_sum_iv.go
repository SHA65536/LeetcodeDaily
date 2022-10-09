package problem0653

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a Binary Search Tree and a target number k,
return true if there exist two elements in the BST such that their sum is equal to the given target.
*/

func findTarget(root *TreeNode, k int) bool {
	var complements = map[int]bool{}

	return findTargetHelper(root, complements, k)
}

func findTargetHelper(root *TreeNode, compl map[int]bool, target int) bool {
	if compl[root.Val] {
		return true
	}
	compl[target-root.Val] = true
	if root.Left != nil && findTargetHelper(root.Left, compl, target) {
		return true
	}
	if root.Right != nil && findTargetHelper(root.Right, compl, target) {
		return true
	}
	return false
}
