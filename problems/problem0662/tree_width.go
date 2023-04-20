package problem0662

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a binary tree, return the maximum width of the given tree.
The maximum width of a tree is the maximum width among all levels.
The width of one level is defined as the length between the end-nodes (the leftmost and rightmost non-null nodes),
where the null nodes between the end-nodes that would be present in a complete binary tree extending down to that level
are also counted into the length calculation.
It is guaranteed that the answer will in the range of a 32-bit signed integer.
*/

type pair struct {
	T *TreeNode
	V int
}

func widthOfBinaryTree(root *TreeNode) int {
	// Res is the max width
	var res int = 1
	// Queue of pair for bfs
	var queue = []pair{{root, 0}}
	for len(queue) > 0 {
		var cnt = len(queue)
		// Start and end are the opposite sides of each level
		var start, end = peek(queue)
		// So the width is the difference between them
		res = max(res, end.V-start.V+1)
		for i := 0; i < cnt; i++ {
			var p, _ = peek(queue)
			var idx = p.V - start.V
			queue, _ = pop(queue)
			if p.T.Left != nil {
				queue = push(queue, pair{p.T.Left, 2*idx + 1})
			}
			if p.T.Right != nil {
				queue = push(queue, pair{p.T.Right, 2*idx + 2})
			}
		}
	}
	return res
}

// pushes a pair value to the queue
func push(queue []pair, val pair) []pair {
	return append(queue, val)
}

// gets the first element and returns the updated queue
func pop(queue []pair) ([]pair, pair) {
	return queue[1:], queue[0]
}

// returns the first and last element
func peek(queue []pair) (pair, pair) {
	return queue[0], queue[len(queue)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
