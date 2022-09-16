package problem0814

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
	{&TreeNode{Val: 1,
		Right: &TreeNode{Val: 0,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 1}}},
		&TreeNode{Val: 1,
			Right: &TreeNode{Val: 0,
				Right: &TreeNode{Val: 1}}}},
}

func TestBinaryTreePruning(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := pruneTree(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
