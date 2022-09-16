package problem0108

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	. "leetcodedaily/helpers/binarytree"
)

type Result struct {
	Input    []int
	Expected *TreeNode
}

var Results = []Result{
	{
		[]int{-10, -3, 0, 5, 9},
		&TreeNode{Val: 0,
			Left: &TreeNode{Val: -3,
				Left: &TreeNode{Val: -10}},
			Right: &TreeNode{Val: 9,
				Left: &TreeNode{Val: 5}}},
	},
}

func TestSortedArrayToBST(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := sortedArrayToBST(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
