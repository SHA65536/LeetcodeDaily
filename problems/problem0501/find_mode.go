package problem0501

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a binary search tree (BST) with duplicates, return all the mode(s) (i.e., the most frequently occurred element) in it.
If the tree has more than one mode, return them in any order.
Assume a BST is defined as follows:
The left subtree of a node contains only nodes with keys less than or equal to the node's key.
The right subtree of a node contains only nodes with keys greater than or equal to the node's key.
Both the left and right subtrees must also be binary search trees.
*/

func findMode(root *TreeNode) []int {
	var res = []int{}
	var freq = map[int]int{}
	var maxFreq int
	getFreqs(freq, &maxFreq, root)
	for i := range freq {
		if freq[i] == maxFreq {
			res = append(res, i)
		}
	}
	return res
}

func getFreqs(freqs map[int]int, maxV *int, root *TreeNode) {
	freqs[root.Val]++
	if freqs[root.Val] > *maxV {
		*maxV = freqs[root.Val]
	}
	if root.Left != nil {
		getFreqs(freqs, maxV, root.Left)
	}
	if root.Right != nil {
		getFreqs(freqs, maxV, root.Right)
	}
}
