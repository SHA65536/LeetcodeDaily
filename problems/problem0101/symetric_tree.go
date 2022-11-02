package problem0101

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
*/

func isSymmetric(root *TreeNode) bool {
	var cur, next []*TreeNode
	cur = []*TreeNode{root}
	for len(cur) > 0 {
		next = make([]*TreeNode, 0, len(cur)*2)
		for i := range cur {
			if cur[i].Val != cur[len(cur)-i-1].Val ||
				((cur[i].Left == nil) != (cur[len(cur)-i-1].Right == nil)) ||
				((cur[i].Right == nil) != (cur[len(cur)-i-1].Left == nil)) {
				return false
			}
			if cur[i].Left != nil {
				next = append(next, cur[i].Left)
			}
			if cur[i].Right != nil {
				next = append(next, cur[i].Right)
			}
		}
		cur = next
	}
	return true
}
