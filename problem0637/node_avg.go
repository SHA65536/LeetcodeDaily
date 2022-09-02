package problem0637

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a binary tree, return the average value of the nodes on each level in the form of an array.
Answers within 10-5 of the actual answer will be accepted.
*/

func averageOfLevels(root *TreeNode) []float64 {
	var results = []float64{}
	var cur = []*TreeNode{root}
	for len(cur) > 0 {
		var sum float64
		var temp = []*TreeNode{}
		for _, node := range cur {
			sum += float64(node.Val)
			if node.Left != nil {
				temp = append(temp, node.Left)
			}
			if node.Right != nil {
				temp = append(temp, node.Right)
			}
		}
		results = append(results, sum/float64(len(cur)))
		cur = temp
	}
	return results
}
