package problem1161

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a binary tree, the level of its root is 1, the level of its children is 2, and so on.
Return the smallest level x such that the sum of all the values of nodes at level x is maximal.
*/

func maxLevelSum(root *TreeNode) int {
	var maxLevel, maxSum, curSum, curLevel int
	var curNodes, nextNodes []*TreeNode
	maxLevel, curLevel = 1, 1
	maxSum = root.Val
	curNodes = []*TreeNode{root}
	for len(curNodes) > 0 {
		curSum = 0
		nextNodes = []*TreeNode{}
		for i := range curNodes {
			curSum += curNodes[i].Val
			if curNodes[i].Left != nil {
				nextNodes = append(nextNodes, curNodes[i].Left)
			}
			if curNodes[i].Right != nil {
				nextNodes = append(nextNodes, curNodes[i].Right)
			}
		}
		if curSum > maxSum {
			maxSum = curSum
			maxLevel = curLevel
		}
		curLevel++
		curNodes = nextNodes
	}
	return maxLevel
}
