package problem0429

import . "leetcodedaily/helpers/narytree"

/*
Given an n-ary tree, return the level order traversal of its nodes' values.
Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value (See examples).
*/

func levelOrder(root *Node) [][]int {
	var results = [][]int{}
	var res []int
	var cur, temp []*Node
	if root != nil {
		cur = []*Node{root}
	}
	for len(cur) > 0 {
		temp = []*Node{}
		res = make([]int, len(cur))
		for i, n := range cur {
			res[i] = n.Val
			temp = append(temp, n.Children...)
		}
		results = append(results, res)
		cur = temp
	}
	return results
}
