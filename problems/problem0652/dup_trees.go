package problem0652

import (
	. "leetcodedaily/helpers/binarytree"
	"strconv"
	"strings"
)

/*
Given the root of a binary tree, return all duplicate subtrees.
For each kind of duplicate subtrees, you only need to return the root node of any one of them.
Two trees are duplicate if they have the same structure with the same node values.
*/

// The idea here is to seralize the subtrees as we go and check for duplicates
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var res []*TreeNode
	var seen = map[string]bool{}
	var dfs func(head *TreeNode) string

	dfs = func(head *TreeNode) string {
		var str strings.Builder
		var rstr string
		if head == nil {
			return ",NIL"
		}
		// Serialize the tree
		str.WriteByte(',')
		str.WriteString(strconv.Itoa(head.Val))
		str.WriteString(dfs(head.Left))
		str.WriteString(dfs(head.Right))

		rstr = str.String()

		// If we've seen the subtree before
		if done, ok := seen[rstr]; ok {
			// Make sure to only include one of each tree
			if !done {
				res = append(res, head)
				seen[rstr] = true
			}
		} else {
			seen[rstr] = false
		}

		return rstr
	}

	dfs(root)
	return res
}
