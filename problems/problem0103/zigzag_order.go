package problem0103

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a binary tree, return the zigzag level order traversal of its nodes' values.
(i.e., from left to right, then right to left for the next level and alternate between).
*/

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res = [][]int{}
	var vals []int
	var cur, next []*TreeNode
	var levelNum int
	if root == nil {
		return res
	}
	cur = []*TreeNode{root}
	for levelNum = 0; len(cur) > 0; levelNum++ {
		next = []*TreeNode{}
		vals = make([]int, len(cur))
		for i := 0; i < len(cur); i++ {
			if levelNum%2 == 0 {
				vals[i] = cur[i].Val
			} else {
				vals[len(cur)-1-i] = cur[i].Val
			}
			if cur[i].Left != nil {
				next = append(next, cur[i].Left)
			}
			if cur[i].Right != nil {
				next = append(next, cur[i].Right)
			}
		}
		cur = next
		res = append(res, vals)
	}
	return res
}
