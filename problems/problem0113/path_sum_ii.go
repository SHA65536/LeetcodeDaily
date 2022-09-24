package problem0113

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a binary tree and an integer targetSum,
return all root-to-leaf paths where the sum of the node values in the path equals targetSum.
Each path should be returned as a list of the node values, not node references.
A root-to-leaf path is a path starting from the root and ending at any leaf node. A leaf is a node with no children.
*/

func pathSum(root *TreeNode, targetSum int) [][]int {
	var res = [][]int{}
	if root != nil {
		solvePath(root, &res, targetSum, 0, []int{})
	}
	return res
}

func solvePath(root *TreeNode, res *[][]int, targetSum, curSum int, path []int) {
	var newPath = make([]int, len(path)+1)
	copy(newPath, path)
	newPath[len(path)] = root.Val
	if root.Left == nil && root.Right == nil {
		if curSum+root.Val == targetSum {
			*res = append(*res, newPath)
		}
	}
	if root.Left != nil {
		solvePath(root.Left, res, targetSum, curSum+root.Val, newPath)
	}
	if root.Right != nil {
		solvePath(root.Right, res, targetSum, curSum+root.Val, newPath)
	}
}
