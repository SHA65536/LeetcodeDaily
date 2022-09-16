package problem1302

import . "leetcodedaily/helpers/binarytree"

/*
https://leetcode.com/problems/deepest-leaves-sum

Given the root of a binary tree, return the sum of values of its deepest leaves.
*/

func deepestLeavesSum(root *TreeNode) int {
	var res int
	var cur = []*TreeNode{root}
	for len(cur) > 0 {
		var next = []*TreeNode{}
		res = 0
		for _, node := range cur {
			res += node.Val
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		cur = next
	}
	return res
}

func deepestLeavesSumTwo(root *TreeNode) int {
	var res int
	var cur, next []*TreeNode
	cur = []*TreeNode{root}
	for len(cur) > 0 {
		next = []*TreeNode{}
		res = 0
		for idx := range cur {
			res += cur[idx].Val
			if cur[idx].Left != nil {
				next = append(next, cur[idx].Left)
			}
			if cur[idx].Right != nil {
				next = append(next, cur[idx].Right)
			}
		}
		cur = next
	}
	return res
}
