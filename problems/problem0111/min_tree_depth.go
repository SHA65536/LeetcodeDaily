package problem0111

import . "leetcodedaily/helpers/binarytree"

/*
Given a binary tree, find its minimum depth.
The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
Note: A leaf is a node with no children.
*/

func minDepth(root *TreeNode) int {
	var res int
	if root == nil {
		return 0
	}
	var cur, nxt []*TreeNode
	cur = []*TreeNode{root}
	for len(cur) > 0 {
		res++
		for _, c := range cur {
			if c.Left == nil && c.Right == nil {
				return res
			}
			if c.Left != nil {
				nxt = append(nxt, c.Left)
			}
			if c.Right != nil {
				nxt = append(nxt, c.Right)
			}
		}
		cur, nxt = nxt, cur[:0]
	}
	return res
}
