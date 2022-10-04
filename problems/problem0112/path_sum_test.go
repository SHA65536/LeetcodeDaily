package problem0112

import (
	"fmt"
	. "leetcodedaily/helpers/binarytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *TreeNode
	Target   int
	Expected bool
}

var Results = []Result{
	{MakeTree(5, 4, 11, 7, NULL, NULL, 2, NULL, NULL, NULL, 8, 13, NULL, NULL, 4, NULL, 1, NULL, NULL), 22, true},
	{MakeTree(1, 2, NULL, NULL, 3, NULL, NULL), 5, false},
	{MakeTree(), 0, false},
}

func TestPathSum(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := hasPathSum(res.Input, res.Target)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
