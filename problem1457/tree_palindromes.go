package problem1457

import . "leetcodedaily/helpers/binarytree"

/*
Given a binary tree where node values are digits from 1 to 9.
A path in the binary tree is said to be pseudo-palindromic if at least one permutation of the node values in the path is a palindrome.
Return the number of pseudo-palindromic paths going from the root node to leaf nodes.
*/

func pseudoPalindromicPaths(root *TreeNode) int {
	var res int
	// Slice representing which numbers in the current path
	// have been seen even or odd times
	var odds = make([]bool, 9)
	findPali(root, &res, odds, 0)
	return res
}

func findPali(root *TreeNode, res *int, odds []bool, oddCnt int) {
	// Updating the count of odd numbers
	if odds[root.Val-1] {
		oddCnt--
	} else {
		oddCnt++
	}
	// Updating the slice
	odds[root.Val-1] = !odds[root.Val-1]

	// Search children nodes
	if root.Left != nil {
		findPali(root.Left, res, odds, oddCnt)
	}
	if root.Right != nil {
		findPali(root.Right, res, odds, oddCnt)
	}

	// If this node is a leaf
	if root.Left == nil && root.Right == nil {
		// The current path is a palindrome if we have at most
		// 1 odd number
		if oddCnt <= 1 {
			*res++
		}
	}

	// When going up, remember to return the slice to the
	// original state
	odds[root.Val-1] = !odds[root.Val-1]
}
