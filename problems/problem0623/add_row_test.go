package problem0623

import (
	"fmt"
	. "leetcodedaily/helpers/binarytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input      *TreeNode
	Val, Depth int
	Expected   *TreeNode
}

var Results = []Result{
	{
		MakeTree(4, 2, 3, NULL, NULL, 1, NULL, NULL, 6, 5, NULL, NULL, NULL),
		1, 2,
		MakeTree(4, 1, 2, 3, NULL, NULL, 1, NULL, NULL, NULL, 1, NULL, 6, 5, NULL, NULL, NULL),
	},
	{
		MakeTree(4, 2, 3, NULL, NULL, 1, NULL, NULL, NULL),
		1, 3,
		MakeTree(4, 2, 1, 3, NULL, NULL, NULL, 1, NULL, 1, NULL, NULL),
	},
}

func TestAddOneRow(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := addOneRow(res.Input, res.Val, res.Depth)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
