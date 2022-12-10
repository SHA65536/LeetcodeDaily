package problem1339

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a binary tree, split the binary tree into two subtrees by removing one edge such that the product of the sums of the subtrees is maximized.
Return the maximum product of the sums of the two subtrees.
Since the answer may be too large, return it modulo 10^9 + 7.
Note that you need to maximize the answer before taking the mod and not after taking it.
*/

const mod uint64 = 1000000007

func maxProduct(root *TreeNode) int {
	var max, total, prod uint64
	var vals []uint64                   // Post order sums
	total = totalPostOrder(root, &vals) // Sum of root
	for i := range vals {
		// vals[i] is sum of the current subtree
		// (total - vals[i]) is the sum of the rest of the tree, excluding current subtree
		prod = vals[i] * (total - vals[i])
		// pick max
		if prod > max {
			max = prod
		}
	}
	return int(max % mod)
}

// totalPostOrder calculates sum of all nodes below and including root
// saves results for each node in a list
func totalPostOrder(root *TreeNode, list *[]uint64) uint64 {
	var sum uint64
	if root == nil {
		return 0
	}
	// Sum of left child
	sum += totalPostOrder(root.Left, list)
	// Sum of right child
	sum += totalPostOrder(root.Right, list)
	// Add self
	sum += uint64(root.Val)
	// Save to list
	(*list) = append((*list), sum)
	return sum
}
