package problem1123

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
	{MakeTree(3, 5, 6, NULL, NULL, 2, 7, NULL, NULL, 4, NULL, NULL, 1, 0, NULL, NULL, 8, NULL, NULL), 2},
	{MakeTree(1), 1},
	{MakeTree(0, 1, NULL, 2, NULL, NULL, 3, NULL, NULL), 2},
}

func TestLowestAncestoyrOfDeepestLeaves(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := lcaDeepestLeaves(res.Input)
		assert.Equal(want, got.Val, fmt.Sprintf("%+v", res))
	}
}
