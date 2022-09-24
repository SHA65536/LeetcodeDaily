package binarytree

const NULL = "#"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MakeTree(vals ...any) *TreeNode {
	var iter int
	return deserialize(vals, &iter)
}

func deserialize(data []any, iter *int) *TreeNode {
	if *iter >= len(data) {
		return nil
	}
	if val, ok := data[*iter].(int); ok {
		*iter++
		return &TreeNode{
			Val:   val,
			Left:  deserialize(data, iter),
			Right: deserialize(data, iter),
		}
	} else {
		*iter++
		return nil
	}
}

func TreeToAny(root *TreeNode) []any {
	var res = make([]any, 0)
	serializeHelper(root, &res)
	return res
}

func serializeHelper(node *TreeNode, res *[]any) {
	if node != nil {
		*res = append(*res, node.Val)
		serializeHelper(node.Left, res)
		serializeHelper(node.Right, res)
	} else {
		*res = append(*res, NULL)
	}
}
