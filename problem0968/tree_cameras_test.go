package problem0968

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
	{&TreeNode{
		Left: &TreeNode{
			Left:  &TreeNode{},
			Right: &TreeNode{},
		},
	}, 1},
	{&TreeNode{
		Left: &TreeNode{
			Left: &TreeNode{
				Left: &TreeNode{
					Left: &TreeNode{},
				},
			},
		},
	}, 2},
}

func TestMinCameraCover(t *testing.T) {
	assert := assert.New(t)
	for _, res := range Results {
		want := res.Expected
		got := minCameraCover(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
