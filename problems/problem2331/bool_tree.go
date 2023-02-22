package problem2331

import . "leetcodedaily/helpers/binarytree"

/*
You are given the root of a full binary tree with the following properties:
    Leaf nodes have either the value 0 or 1, where 0 represents False and 1 represents True.
    Non-leaf nodes have either the value 2 or 3, where 2 represents the boolean OR and 3 represents the boolean AND.
The evaluation of a node is as follows:
    If the node is a leaf node, the evaluation is the value of the node, i.e. True or False.
    Otherwise, evaluate the node's two children and apply the boolean operation of its value with the children's evaluations.
Return the boolean result of evaluating the root node.
A full binary tree is a binary tree where each node has either 0 or 2 children.
A leaf node is a node that has zero children.
*/

const (
	FALSE = iota
	TRUE
	OR
	AND
)

func evaluateTree(root *TreeNode) bool {
	switch root.Val {
	case FALSE:
		return false
	case TRUE:
		return true
	case OR:
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	case AND:
		return evaluateTree(root.Left) && evaluateTree(root.Right)
	}
	return false
}
