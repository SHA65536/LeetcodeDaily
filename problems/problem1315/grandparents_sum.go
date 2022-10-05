package problem1315

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a binary tree, return the sum of values of nodes with an even-valued grandparent.
If there are no nodes with an even-valued grandparent, return 0.
A grandparent of a node is the parent of its parent if it exists.
*/

func sumEvenGrandparent(root *TreeNode) int {
	var res int
	evenGrandparents(root, &res, -1, -1)
	return res
}

func evenGrandparents(root *TreeNode, res *int, grand, parent int) {
	if grand%2 == 0 {
		*res += root.Val
	}
	if root.Left != nil {
		evenGrandparents(root.Left, res, parent, root.Val)
	}
	if root.Right != nil {
		evenGrandparents(root.Right, res, parent, root.Val)
	}
}
