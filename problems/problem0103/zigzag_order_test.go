package problem0103

import (
	"fmt"
	. "leetcodedaily/helpers/binarytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *TreeNode
	Expected [][]int
}

var Results = []Result{
	{MakeTree(3, 9, NULL, NULL, 20, 15, NULL, NULL, 7, NULL, NULL), [][]int{{3}, {20, 9}, {15, 7}}},
	{MakeTree(1), [][]int{{1}}},
	{MakeTree(), [][]int{}},
}

func TestZigZagOrder(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := zigzagLevelOrder(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
