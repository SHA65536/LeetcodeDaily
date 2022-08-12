package problem0235

import (
	"fmt"
	"testing"

	. "leetcodedaily/helpers/binarytree"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *TreeNode
	P, Q     *TreeNode
	Expected *TreeNode
}

var Tree = &TreeNode{Val: 6,
	Left: &TreeNode{Val: 2,
		Left: &TreeNode{Val: 0},
		Right: &TreeNode{Val: 4,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 5}}},
	Right: &TreeNode{Val: 8,
		Left:  &TreeNode{Val: 7},
		Right: &TreeNode{Val: 9}}}

var Results = []Result{
	{
		Tree,
		&TreeNode{Val: 2}, &TreeNode{Val: 8},
		Tree,
	},
	{
		Tree,
		&TreeNode{Val: 2}, &TreeNode{Val: 4},
		Tree.Left,
	},
	{
		Tree,
		&TreeNode{Val: 3}, &TreeNode{Val: 5},
		Tree.Left.Right,
	},
}

func TestLowestAncestorBinSearch(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := lowestCommonAncestor(res.Input, res.P, res.Q)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
