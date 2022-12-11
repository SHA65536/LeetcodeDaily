package problem0124

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
	{MakeTree(1, 2, NULL, NULL, 3, NULL, NULL), 6},
	{MakeTree(-10, 9, NULL, NULL, 20, 15, NULL, NULL, 7, NULL, NULL), 42},
}

func TestMaxSplitTree(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maxPathSum(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
