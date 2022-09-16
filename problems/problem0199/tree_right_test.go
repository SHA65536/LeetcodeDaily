package problem0199

import (
	"fmt"
	"testing"

	. "leetcodedaily/helpers/binarytree"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *TreeNode
	Expected []int
}

var Results = []Result{
	{&TreeNode{Val: 1,
		Right: &TreeNode{Val: 3,
			Right: &TreeNode{Val: 4}},
		Left: &TreeNode{Val: 2,
			Right: &TreeNode{Val: 5}}},
		[]int{1, 3, 4}},
	{&TreeNode{Val: 1,
		Right: &TreeNode{Val: 3}},
		[]int{1, 3}},
	{&TreeNode{Val: 1,
		Left: &TreeNode{Val: 3}},
		[]int{1, 3}},
}

func TestRightSideView(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := rightSideView(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
