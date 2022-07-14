package problem0105

import (
	"fmt"
	"testing"

	. "leetcodedaily/helpers/binarytree"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	PreOrder []int
	InOrder  []int
	Expected *TreeNode
}

var Results = []Result{
	{
		[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7},
		&TreeNode{Val: 3,
			Left: &TreeNode{Val: 9},
			Right: &TreeNode{Val: 20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7}}},
	},
	{
		[]int{-1}, []int{-1},
		&TreeNode{Val: -1},
	},
}

func TestBuildTree(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := buildTree(res.PreOrder, res.InOrder)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
