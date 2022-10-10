package problem0098

import (
	"fmt"
	"testing"

	. "leetcodedaily/helpers/binarytree"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *TreeNode
	Expected bool
}

var Results = []Result{
	{
		&TreeNode{Val: 2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3}},
		true,
	},
	{
		&TreeNode{Val: 2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 0}},
		false,
	},
	{
		&TreeNode{Val: 5,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{Val: 4,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 6}}},
		false,
	},
	{
		&TreeNode{Val: 5,
			Left: &TreeNode{Val: 4},
			Right: &TreeNode{Val: 7,
				Left:  &TreeNode{Val: 6},
				Right: &TreeNode{Val: 8}}},
		true,
	},
}

func TestValidateBST(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := isValidBST(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
