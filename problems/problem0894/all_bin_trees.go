package problem0894

import . "leetcodedaily/helpers/binarytree"

/*
Given an integer n, return a list of all possible full binary trees with n nodes.
Each node of each tree in the answer must have Node.val == 0.
Each element of the answer is the root node of one possible tree.
You may return the final list of trees in any order.
A full binary tree is a binary tree where each node has exactly 0 or 2 children.
*/

func allPossibleFBT(n int) []*TreeNode {
	return helper(make([][]*TreeNode, n+1), n)
}

func helper(dp [][]*TreeNode, n int) []*TreeNode {
	if dp[n] != nil {
		return dp[n]
	}
	var res = []*TreeNode{}
	if n == 1 {
		res = append(res, &TreeNode{})
	} else if n%2 != 0 {
		for i := 1; i < n; i += 2 {
			for _, l := range helper(dp, i) {
				for _, r := range helper(dp, n-i-1) {
					res = append(res, &TreeNode{Val: 0, Left: l, Right: r})
				}
			}
		}
	}
	dp[n] = res
	return res
}
