package problem0102

import . "leetcodedaily/helpers/binarytree"

/*
https://leetcode.com/problems/binary-tree-level-order-traversal/

Given the root of a binary tree, return the level order traversal of its nodes' values.
(i.e., from left to right, level by level).
*/

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int = make([][]int, 0)
	var cur = []*TreeNode{root}
	for len(cur) > 0 {
		var next = make([]*TreeNode, 0)
		var vals = make([]int, len(cur))
		for idx, node := range cur {
			vals[idx] = node.Val
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		cur = next
		res = append(res, vals)
	}
	return res
}
