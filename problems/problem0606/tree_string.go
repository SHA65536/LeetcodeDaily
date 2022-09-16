package problem0606

import (
	"fmt"
	. "leetcodedaily/helpers/binarytree"
	"strconv"
)

/*
Given the root of a binary tree, construct a string consisting of parenthesis and integers from a binary tree with the preorder traversal way, and return it.
Omit all the empty parenthesis pairs that do not affect the one-to-one mapping relationship between the string and the original binary tree.
*/

func tree2str(root *TreeNode) string {
	if root.Left == nil && root.Right == nil {
		return strconv.Itoa(root.Val)
	}
	if root.Left == nil {
		return fmt.Sprintf("%d()(%s)", root.Val, tree2str(root.Right))
	}
	if root.Right == nil {
		return fmt.Sprintf("%d(%s)", root.Val, tree2str(root.Left))
	}
	return fmt.Sprintf("%d(%s)(%s)", root.Val, tree2str(root.Left), tree2str(root.Right))
}
