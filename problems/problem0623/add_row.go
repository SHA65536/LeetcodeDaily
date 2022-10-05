package problem0623

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a binary tree and two integers val and depth, add a row of nodes with value val at the given depth depth.
Note that the root node is at depth 1.
The adding rule is:
Given the integer depth, for each not null tree node cur at the depth depth - 1,
create two tree nodes with value val as cur's left subtree root and right subtree root.
cur's original left subtree should be the left subtree of the new left subtree root.
cur's original right subtree should be the right subtree of the new right subtree root.
If depth == 1 that means there is no depth depth - 1 at all,
then create a tree node with value val as the new root of the whole original tree, and the original tree is the new root's left subtree.
*/

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	var cur, next []*TreeNode
	// Special case if target is above root
	if depth == 1 {
		return &TreeNode{Val: val, Left: root}
	}
	cur = []*TreeNode{root}
	// Finding row above target depth
	for depth > 2 {
		next = []*TreeNode{}
		for _, node := range cur {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		cur = next
		depth--
	}
	// Switching nodes of target
	for _, node := range cur {
		node.Left = &TreeNode{
			Val:  val,
			Left: node.Left,
		}
		node.Right = &TreeNode{
			Val:   val,
			Right: node.Right,
		}
	}
	return root
}
