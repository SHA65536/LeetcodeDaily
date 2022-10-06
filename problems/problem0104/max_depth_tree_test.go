package problem0104

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
	{MakeTree(3, 9, NULL, NULL, 20, 25, NULL, NULL, 7, NULL, NULL), 3},
	{MakeTree(1, NULL, 2, NULL, NULL), 2},
	{nil, 0},
	{MakeTree(1), 1},
}

func TestMaxTreeDepth(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maxDepth(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
