package problem0094

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
		Right: &TreeNode{Val: 2,
			Left: &TreeNode{Val: 3}}},
		[]int{1, 3, 2}},
}

func TestInorderTree(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := inorderTraversal(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
