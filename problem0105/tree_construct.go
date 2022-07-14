package problem0105

import (
	. "leetcodedaily/helpers/binarytree"
)

/*
https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal

Given two integer arrays `preorder` and `inorder`
Where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree,
construct and return the binary tree.
*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	var stack []*TreeNode = make([]*TreeNode, 0)
	var root *TreeNode = &TreeNode{Val: preorder[0]}
	var cur = root
	for i, j := 1, 0; i < len(preorder); i++ {
		if cur.Val != inorder[j] {
			cur.Left = &TreeNode{Val: preorder[i]}
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			j++
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[j] {
				cur = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				j++
			}
			cur.Right = &TreeNode{Val: preorder[i]}
			cur = cur.Right
		}
	}
	return root
}
