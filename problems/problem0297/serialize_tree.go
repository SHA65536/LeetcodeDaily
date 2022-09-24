package problem0297

import (
	. "leetcodedaily/helpers/binarytree"
	"strconv"
	"strings"
)

/*
Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer,
or transmitted across a network connection link to be reconstructed later in the same or another computer environment.
Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work.
You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.
*/

const SPLIT = ","
const NULL = "N"

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var res string
	serializeHelper(root, &res)
	return res
}

func serializeHelper(node *TreeNode, res *string) {
	if node == nil {
		*res += NULL + SPLIT
	} else {
		*res += strconv.Itoa(node.Val) + SPLIT
		serializeHelper(node.Left, res)
		serializeHelper(node.Right, res)
	}
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, SPLIT)
	nodes := make(chan string, len(vals))
	for i := range vals {
		nodes <- vals[i]
	}
	return deserializeHelper(nodes)
}

func deserializeHelper(nodes chan string) *TreeNode {
	var root *TreeNode
	var val = <-nodes
	if val != NULL {
		num, err := strconv.Atoi(val)
		if err != nil {
			return nil
		}
		root = &TreeNode{
			Val:   num,
			Left:  deserializeHelper(nodes),
			Right: deserializeHelper(nodes),
		}
	}
	return root
}
