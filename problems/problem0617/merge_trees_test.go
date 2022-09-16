package problem0617

import (
	"fmt"
	. "leetcodedaily/helpers/binarytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Tree1    *TreeNode
	Tree2    *TreeNode
	Expected *TreeNode
}

var Results = []Result{
	{
		Tree1: &TreeNode{Val: 1,
			Left: &TreeNode{Val: 3,
				Left: &TreeNode{Val: 5}},
			Right: &TreeNode{Val: 2}},
		Tree2: &TreeNode{Val: 2,
			Left: &TreeNode{Val: 1,
				Right: &TreeNode{Val: 4}},
			Right: &TreeNode{Val: 3,
				Right: &TreeNode{Val: 7}}},
		Expected: &TreeNode{Val: 3,
			Left: &TreeNode{Val: 4,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 4}},
			Right: &TreeNode{Val: 5,
				Right: &TreeNode{Val: 7}}},
	},
}

func TestMergeTrees(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := mergeTrees(res.Tree1, res.Tree2)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
