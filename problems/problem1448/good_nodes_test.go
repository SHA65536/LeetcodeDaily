package problem1448

import (
	"fmt"
	"testing"

	. "leetcodedaily/helpers/binarytree"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *TreeNode
	Expected int
}

var Results = []Result{
	{&TreeNode{Val: 3,
		Left: &TreeNode{Val: 1,
			Left: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 5}}}, 4},
	{&TreeNode{Val: 3,
		Left: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 2}}}, 3},
}

func TestGoodNodes(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := goodNodes(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
