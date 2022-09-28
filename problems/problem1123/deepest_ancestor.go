package problem1123

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a binary tree, return the lowest common ancestor of its deepest leaves.
Recall that:
The node of a binary tree is a leaf if and only if it has no children
The depth of the root of the tree is 0. if the depth of a node is d, the depth of each of its children is d + 1.
The lowest common ancestor of a set S of nodes, is the node A with the largest depth such that every node in S is in the subtree with root A.
*/

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	var depth = map[*TreeNode]int{}
	return findLca(depth, root)
}

func findLca(depth map[*TreeNode]int, root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var l, r int = findDepth(depth, root.Left), findDepth(depth, root.Right)
	if l == r {
		return root
	}
	if l > r {
		return findLca(depth, root.Left)
	} else {
		return findLca(depth, root.Right)
	}
}

func findDepth(depth map[*TreeNode]int, root *TreeNode) int {
	if root == nil {
		return 0
	}
	if val, ok := depth[root]; ok {
		return val
	} else {
		l, r := findDepth(depth, root.Left), findDepth(depth, root.Right)
		if l > r {
			depth[root] = 1 + l
		} else {
			depth[root] = 1 + r
		}
		return depth[root]
	}
}
