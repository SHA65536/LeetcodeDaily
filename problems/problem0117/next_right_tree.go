package problem0117

import . "leetcodedaily/helpers/binarytree"

/*
https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii

Given a Binary Tree,
Populate each next pointer to point to its next right node.
If there is no next right node, the next pointer should be set to NULL.

Initially, all next pointers are set to NULL.
*/

func connect(root *Node) *Node {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	var cur []*Node
	var last = []*Node{root}
	var prev *Node
	for len(last) > 0 {
		cur = []*Node{}
		for idx, n := range last {
			if idx != 0 {
				prev.Next = n
			}
			prev = n
			if n.Left != nil {
				cur = append(cur, n.Left)
			}
			if n.Right != nil {
				cur = append(cur, n.Right)
			}
		}
		last = cur
	}
	return root
}
