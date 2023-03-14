package problem0129

import (
	. "leetcodedaily/helpers/binarytree"
	"math"
)

/*
You are given the root of a binary tree containing digits from 0 to 9 only.
Each root-to-leaf path in the tree represents a number.
    For example, the root-to-leaf path 1 -> 2 -> 3 represents the number 123.
Return the total sum of all root-to-leaf numbers.
Test cases are generated so that the answer will fit in a 32-bit integer.
A leaf node is a node with no children.
*/

func sumNumbers(root *TreeNode) int {
	var res int
	sumHelper(root, &res)
	return res
}

func sumHelper(root *TreeNode, res *int) []int {
	var depth []int
	// If this is a leaf, add the value to res, and return the next depth of 10
	if root.Left == nil && root.Right == nil {
		*res += root.Val
		return []int{10}
	}
	// We're going bottom to top each depth's value is 10 times bigger
	if root.Left != nil {
		// Get the depths from the bottom
		for _, d := range sumHelper(root.Left, res) {
			// Return depth bigger by 10
			depth = append(depth, d*10)
			// Add val times depth to the res
			*res += root.Val * d
		}
	}
	if root.Right != nil {
		for _, d := range sumHelper(root.Right, res) {
			depth = append(depth, d*10)
			*res += root.Val * d
		}
	}
	return depth
}

func sumNumbersIterative(root *TreeNode) int {
	var sum, cur, depth int
	for root != nil {
		if root.Left != nil {
			var prev = root.Left
			depth = 1
			for prev.Right != nil && prev.Right != root {
				prev = prev.Right
				depth++
			}
			if prev.Right == nil {
				prev.Right = root
				cur = cur*10 + root.Val
				root = root.Left
			} else {
				prev.Right = nil
				if prev.Left == nil {
					sum += cur
				}
				cur /= int(math.Pow10(depth))
				root = root.Right
			}
		} else {
			cur = cur*10 + root.Val
			if root.Right == nil {
				sum += cur
			}
			root = root.Right
		}
	}
	return sum
}
