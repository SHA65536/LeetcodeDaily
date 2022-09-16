package problem0606

import (
	"fmt"
	"testing"

	. "leetcodedaily/helpers/binarytree"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *TreeNode
	Expected string
}

var Results = []Result{
	{&TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 3}}, "1(2(4))(3)"},
	{&TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 3}}, "1(2()(4))(3)"},
	{&TreeNode{Val: 1}, "1"},
}

func TestTreeToString(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := tree2str(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
