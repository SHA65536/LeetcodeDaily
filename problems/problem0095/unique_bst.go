package problem0095

import (
	. "leetcodedaily/helpers/binarytree"
)

/*
Given an integer n, return all the structurally unique BST's (binary search trees),
which has exactly n nodes of unique values from 1 to n.
Return the answer in any order.
*/

func generateTrees(n int) []*TreeNode {
	var calcTrees func(vals []int) []*TreeNode
	var vals = make([]int, n)
	var cache = make([][][]*TreeNode, n)
	for i := range cache {
		cache[i] = make([][]*TreeNode, n)
	}

	// calcTrees returns all possible BST trees for values vals
	calcTrees = func(vals []int) []*TreeNode {
		// No trees can be generated
		if len(vals) == 0 {
			return []*TreeNode{nil}
		}

		// Leaf node
		if len(vals) == 1 {
			return []*TreeNode{{Val: vals[0]}}
		}

		// Check if already calculated
		if res := cache[vals[0]-1][vals[len(vals)-1]-1]; res != nil {
			return res
		}

		var res []*TreeNode
		// Trying all possible divisions
		for i := range vals {
			left := calcTrees(vals[:i])
			right := calcTrees(vals[i+1:])
			// Adding all division combinations to res
			for l := range left {
				for r := range right {
					res = append(res, &TreeNode{
						Val:   vals[i],
						Left:  left[l],
						Right: right[r],
					})
				}
			}
		}

		cache[vals[0]-1][vals[len(vals)-1]-1] = res

		return res
	}

	for i := 1; i <= n; i++ {
		vals[i-1] = i
	}

	return calcTrees(vals)
}
