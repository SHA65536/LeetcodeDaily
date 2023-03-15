package problem0958

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a binary tree, determine if it is a complete binary tree.
In a complete binary tree, every level, except possibly the last, is completely filled,
and all nodes in the last level are as far left as possible. It can have between 1 and 2h nodes inclusive at the last level h.
*/

func isCompleteTree(root *TreeNode) bool {
	var depth int = 1
	var cur, nxt []*TreeNode
	var sawEmpty bool
	cur = []*TreeNode{root}
	for len(cur) > 0 {
		depth *= 2
		nxt = make([]*TreeNode, 0, depth)
		sawEmpty = false
		for _, cn := range cur {
			if cn.Left != nil {
				if sawEmpty {
					return false
				}
				nxt = append(nxt, cn.Left)
			} else {
				sawEmpty = true
			}
			if cn.Right != nil {
				if sawEmpty {
					return false
				}
				nxt = append(nxt, cn.Right)
			} else {
				sawEmpty = true
			}
		}
		cur = nxt
		if len(cur) != depth {
			break
		}
	}
	for _, cn := range cur {
		if cn.Left != nil || cn.Right != nil {
			return false
		}
	}
	return true
}
