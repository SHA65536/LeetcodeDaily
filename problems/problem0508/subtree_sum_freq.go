package problem0508

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a binary tree, return the most frequent subtree sum.
If there is a tie, return all the values with the highest frequency in any order.
The subtree sum of a node is defined as the sum of all the node values formed by the
subtree rooted at that node (including the node itself).
*/

func findFrequentTreeSum(root *TreeNode) []int {
	var maxFreq int
	var res []int
	var freqs = map[int]int{}

	// Get frequencies
	getFreqs(root, freqs)

	// Find top frequency
	for _, v := range freqs {
		if v > maxFreq {
			maxFreq = v
		}
	}

	// Finding sums with top frequency
	for k, v := range freqs {
		if v == maxFreq {
			res = append(res, k)
		}
	}

	return res
}

func getFreqs(root *TreeNode, freqs map[int]int) (sum int) {
	if root == nil {
		return 0
	}
	sum = root.Val + getFreqs(root.Left, freqs) + getFreqs(root.Right, freqs)
	freqs[sum]++
	return sum
}
