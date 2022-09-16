package problem0987

import (
	"fmt"
	"testing"

	. "leetcodedaily/helpers/binarytree"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *TreeNode
	Expected [][]int
}

var Results = []Result{
	{&TreeNode{Val: 3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7}}},
		[][]int{{9}, {3, 15}, {20}, {7}}},
	{&TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7}}},
		[][]int{{4}, {2}, {1, 5, 6}, {3}, {7}}},
	{&TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 6}},
		Right: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 7}}},
		[][]int{{4}, {2}, {1, 5, 6}, {3}, {7}}},
}

func TestAverageOfLevels(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := verticalTraversal(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
