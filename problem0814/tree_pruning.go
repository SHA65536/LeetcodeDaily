package problem0814

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a binary tree, return the same tree where every subtree (of the given tree) not containing a 1 has been removed.
A subtree of a node node is node plus every node that is a descendant of node.
*/

func pruneTree(root *TreeNode) *TreeNode {
	next := &TreeNode{Right: root}
	solveTree(next)
	return next.Right
}

func solveTree(root *TreeNode) int {
	var res int = root.Val
	if root.Left != nil {
		if solveTree(root.Left) == 1 {
			res = 1
		} else {
			root.Left = nil
		}
	}
	if root.Right != nil {
		if solveTree(root.Right) == 1 {
			res = 1
		} else {
			root.Right = nil
		}
	}
	return res
}
