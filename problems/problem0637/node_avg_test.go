package problem0637

import (
	"fmt"
	"testing"

	. "leetcodedaily/helpers/binarytree"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *TreeNode
	Expected []float64
}

var Results = []Result{
	{&TreeNode{Val: 3,
		Left: &TreeNode{Val: 1,
			Left: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 5}}}, []float64{3, 2.5, 3}},
	{&TreeNode{Val: 3,
		Left: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 2}}}, []float64{3, 3, 3}},
}

func TestAverageOfLevels(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := averageOfLevels(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
