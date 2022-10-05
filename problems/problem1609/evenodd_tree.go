package problem1609

import . "leetcodedaily/helpers/binarytree"

/*
A binary tree is named Even-Odd if it meets the following conditions:
The root of the binary tree is at level index 0, its children are at level index 1, their children are at level index 2, etc.
For every even-indexed level, all nodes at the level have odd integer values in strictly increasing order (from left to right).
For every odd-indexed level, all nodes at the level have even integer values in strictly decreasing order (from left to right).
Given the root of a binary tree, return true if the binary tree is Even-Odd, otherwise return false.
*/

const MAXVAL = 1000001
const MINVAL = 0

func isEvenOddTree(root *TreeNode) bool {
	var level, prev int
	var cur, next []*TreeNode
	for cur = []*TreeNode{root}; len(cur) > 0; level++ {
		next = []*TreeNode{}
		if level%2 != 0 {
			prev = MAXVAL
			for _, node := range cur {
				if prev <= node.Val || node.Val%2 != 0 {
					return false
				}
				if node.Left != nil {
					next = append(next, node.Left)
				}
				if node.Right != nil {
					next = append(next, node.Right)
				}
				prev = node.Val
			}
		} else {
			prev = MINVAL
			for _, node := range cur {
				if prev >= node.Val || node.Val%2 == 0 {
					return false
				}
				if node.Left != nil {
					next = append(next, node.Left)
				}
				if node.Right != nil {
					next = append(next, node.Right)
				}
				prev = node.Val
			}
		}
		cur = next
	}
	return true
}
