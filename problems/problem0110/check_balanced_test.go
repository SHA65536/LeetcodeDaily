package problem0110

import (
	"fmt"
	. "leetcodedaily/helpers/binarytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *TreeNode
	Expected bool
}

var Results = []Result{
	{MakeTree(3, 8, NULL, NULL, 20, 15, NULL, NULL, 7, NULL, NULL), true},
	{MakeTree(1, 2, 3, 4, NULL, NULL, 4, NULL, NULL, 3, NULL, NULL, 2, NULL, NULL), false},
}

func TestCheckTreeBalance(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := isBalanced(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
