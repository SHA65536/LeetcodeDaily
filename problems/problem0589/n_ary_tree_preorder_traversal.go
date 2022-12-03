package problem0589

import . "leetcodedaily/helpers/narytree"

/*
Given the root of an n-ary tree, return the preorder traversal of its nodes' values.
Nary-Tree input serialization is represented in their level order traversal.
Each group of children is separated by the null value (See examples)
*/

func preorder(root *Node) []int {
	var res = []int{}
	if root == nil {
		return res
	}
	// Appending current value
	res = append(res, root.Val)
	if root.Children != nil {
		// Recursively adding the children to result
		for i := range root.Children {
			res = append(res, preorder(root.Children[i])...)
		}
	}
	return res
}

func preorderShared(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}
	// Same as preorder but with only one list
	visitChildren(root, &res)
	return res
}

func visitChildren(root *Node, res *[]int) {
	// Appending current value
	(*res) = append(*res, root.Val)
	if root.Children != nil {
		// Recursively adding the children to result
		for i := range root.Children {
			visitChildren(root.Children[i], res)
		}
	}
}

func preorderStack(root *Node) []int {
	var res []int
	var stack = NodeStack{root}
	if root == nil {
		return res
	}
	// Looping until stack is empty
	for len(stack) > 0 {
		// Adding top value to result
		cur := stack.Pop()
		res = append(res, cur.Val)
		if cur.Children != nil {
			// Adding all the children in reverse order to the stack
			for i := len(cur.Children) - 1; i >= 0; i-- {
				stack.Push(cur.Children[i])
			}
		}
	}
	return res
}

type NodeStack []*Node

func (s *NodeStack) Pop() *Node {
	ret := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return ret
}

func (s *NodeStack) Push(val *Node) {
	(*s) = append((*s), val)
}
