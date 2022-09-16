package problem0114

import (
	"fmt"
	"testing"

	. "leetcodedaily/helpers/binarytree"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *TreeNode
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
		&TreeNode{Val: 3, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 7, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 1, Right: &TreeNode{Val: 0, Right: &TreeNode{Val: 8}}}}}}}}}},
	{&TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
		&TreeNode{Val: 1, Right: &TreeNode{Val: 2}}},
	{&TreeNode{Val: -1,
		Left: &TreeNode{Val: 0,
			Left: &TreeNode{Val: -2,
				Left: &TreeNode{Val: 8}},
			Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 3}},
		&TreeNode{Val: -1, Right: &TreeNode{Val: 0, Right: &TreeNode{Val: -2, Right: &TreeNode{Val: 8, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 3}}}}}}},
}

func TestFlatten(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		flatten(res.Input)
		assert.Equal(want, res.Input, fmt.Sprintf("%+v", res))
	}
}
