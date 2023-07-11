package problem0863

import . "leetcodedaily/helpers/binarytree"

/*
Given the root of a binary tree, the value of a target node target, and an integer k,
return an array of the values of all nodes that have a distance k from the target node.
You can return the answer in any order.
*/

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	var res []int
	// parents[i] is the parent node of i
	var parents = map[*TreeNode]*TreeNode{}
	// visited to see prevent duplicates
	var visited = map[*TreeNode]struct{}{}
	// dfs traverses the graph from head, and records nodes d distance away
	var dfs func(head *TreeNode, d int)
	makeParents(root, parents)

	dfs = func(head *TreeNode, d int) {
		// If we already visited, continue
		if _, ok := visited[head]; ok {
			return
		}
		// Add to set
		visited[head] = struct{}{}
		if d == 0 {
			// If we got to the target, add to res
			res = append(res, head.Val)
			return
		}
		if head.Left != nil {
			// Go down to the left
			dfs(head.Left, d-1)
		}
		if head.Right != nil {
			// Go down to the right
			dfs(head.Right, d-1)
		}
		if par := parents[head]; par != nil {
			// Go up if there's up
			dfs(par, d-1)
		}
	}
	dfs(target, k)
	return res
}

// makeParents traverses the tree and registers the parent node of each node
func makeParents(root *TreeNode, parents map[*TreeNode]*TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		parents[root.Left] = root
		makeParents(root.Left, parents)
	}
	if root.Right != nil {
		parents[root.Right] = root
		makeParents(root.Right, parents)
	}
}
