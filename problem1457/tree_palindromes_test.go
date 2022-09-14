package problem1457

import (
	"fmt"
	. "leetcodedaily/helpers/binarytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *TreeNode
	Expected int
}

var Results = []Result{
	{&TreeNode{Val: 2,
		Left: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 1}},
		Right: &TreeNode{Val: 1,
			Right: &TreeNode{Val: 1}}}, 2},
	{&TreeNode{Val: 2,
		Left: &TreeNode{Val: 1,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3,
				Right: &TreeNode{Val: 1}}},
		Right: &TreeNode{Val: 1}}, 1},
	{&TreeNode{Val: 9}, 1},
}

func TestTreePsuedoPali(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := pseudoPalindromicPaths(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
