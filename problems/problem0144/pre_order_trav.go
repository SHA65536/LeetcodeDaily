package problem0144

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a binary tree, return the preorder traversal of its nodes' values.
*/

func preorderTraversal(root *TreeNode) []int {
	var res = []int{}
	preorderHelper(root, &res)
	return res
}

func preorderHelper(root *TreeNode, res *[]int) {
	if root != nil {
		*res = append(*res, root.Val)
		preorderHelper(root.Left, res)
		preorderHelper(root.Right, res)
	}
}

func preorderTraversalStack(root *TreeNode) []int {
	var res []int
	var stack = []*TreeNode{root}
	if root == nil {
		return []int{}
	}
	for len(stack) > 0 {
		// Get last element
		cur := stack[len(stack)-1]
		// Remove last element from stack
		stack = stack[:len(stack)-1]
		// Add to results
		res = append(res, cur.Val)
		// Try to add right branch
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		// Try to add left branch
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		// It's important to add the right before the left
		// since the stack is LIFO
	}
	return res
}
