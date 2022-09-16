package problem0199

import . "leetcodedaily/helpers/binarytree"

/*
https://leetcode.com/problems/binary-tree-right-side-view

Given the root of a binary tree,
imagine yourself standing on the right side of it,
return the values of the nodes you can see ordered from top to bottom.
*/

// We'll approach this using breadth first search
func rightSideView(root *TreeNode) []int {
	var result []int
	var bfs, nextbfs []*TreeNode
	if root == nil {
		return []int{}
	}
	bfs = []*TreeNode{root}
	for len(bfs) > 0 {
		// Adding right most value to the tree
		result = append(result, bfs[len(bfs)-1].Val)
		nextbfs = []*TreeNode{}
		for _, node := range bfs {
			if node.Left != nil {
				nextbfs = append(nextbfs, node.Left)
			}
			if node.Right != nil {
				nextbfs = append(nextbfs, node.Right)
			}
		}
		bfs = nextbfs
	}
	return result
}
