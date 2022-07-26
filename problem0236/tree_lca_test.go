package problem0236

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

var Results = []Result{
	{&TreeNode{Val: 3,
		Left: &TreeNode{Val: 5,
			Left: &TreeNode{Val: 6},
			Right: &TreeNode{Val: 2,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 4}}},
		Right: &TreeNode{Val: 1,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 8}}},
		&TreeNode{Val: 5}, &TreeNode{Val: 1}, &TreeNode{Val: 3}},
	{&TreeNode{Val: 3,
		Left: &TreeNode{Val: 5,
			Left: &TreeNode{Val: 6},
			Right: &TreeNode{Val: 2,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 4}}},
		Right: &TreeNode{Val: 1,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 8}}},
		&TreeNode{Val: 5}, &TreeNode{Val: 4}, &TreeNode{Val: 5}},
	{&TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
		&TreeNode{Val: 1}, &TreeNode{Val: 2}, &TreeNode{Val: 1}},
	{&TreeNode{Val: -1,
		Left: &TreeNode{Val: 0,
			Left: &TreeNode{Val: -2,
				Left: &TreeNode{Val: 8}},
			Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 3}},
		&TreeNode{Val: 3}, &TreeNode{Val: 8}, &TreeNode{Val: -1}},
}

func TestLowestCommonAncestor(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := lowestCommonAncestor(res.Input, res.P, res.Q)
		assert.Equal(want.Val, got.Val, fmt.Sprintf("%+v", res))
	}
}
