package problem0297

import (
	"fmt"
	. "leetcodedaily/helpers/binarytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Tree *TreeNode
}

var Results = []Result{
	{&TreeNode{Val: 1,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5}}}},
	{&TreeNode{Val: 1,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3,
			Left: &TreeNode{Val: 4}}}},
}

func TestSerializeTree(t *testing.T) {
	assert := assert.New(t)

	ser := Constructor()

	for _, res := range Results {

		out := ser.serialize(res.Tree)
		ret := ser.deserialize(out)
		assert.Equal(ret, res.Tree, fmt.Sprintf("%+v", res))
	}
}
