package problem0515

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a binary tree, return an array of the largest value in each row of the tree (0-indexed).
*/

func largestValues(root *TreeNode) []int {
	var res = []int{}
	var cur, next []*TreeNode
	if root == nil {
		return nil
	}
	cur = []*TreeNode{root}
	for len(cur) > 0 {
		var max = cur[0].Val
		next = []*TreeNode{}
		for _, node := range cur {
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		res = append(res, max)
		cur = next
	}
	return res
}
