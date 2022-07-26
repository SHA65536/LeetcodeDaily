package problem0236

import . "leetcodedaily/helpers/binarytree"

/*
Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the definition of LCA on Wikipedia:
“The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants
(where we allow a node to be a descendant of itself).”
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	_, _, res := lowestCommonDFS(root, p, q, false)
	return res
}

func lowestCommonDFS(root, p, q *TreeNode, halfDone bool) (found, done bool, res *TreeNode) {
	var childFound, childDone bool
	var childRes *TreeNode
	var selfFound = root.Val == p.Val || root.Val == q.Val
	if selfFound && halfDone {
		return true, false, nil
	}
	if root.Left == nil && root.Right == nil {
		return selfFound, false, nil
	}
	if root.Left != nil {
		childFound, childDone, childRes = lowestCommonDFS(root.Left, p, q, selfFound)
	}
	if childDone {
		return true, true, childRes
	}
	if childFound {
		if selfFound {
			return true, true, root
		} else {
			selfFound = true
			childFound = false
		}
	}
	if root.Right != nil {
		childFound, childDone, childRes = lowestCommonDFS(root.Right, p, q, selfFound)
	}
	if childDone {
		return true, true, childRes
	}
	if childFound {
		if selfFound {
			return true, true, root
		} else {
			selfFound = true
		}
	}
	return selfFound, false, nil
}
