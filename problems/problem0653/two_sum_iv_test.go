package problem0653

import (
	"fmt"
	"testing"

	. "leetcodedaily/helpers/binarytree"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *TreeNode
	Target   int
	Expected bool
}

var Results = []Result{
	{MakeTree(5, 3, 2, NULL, NULL, 4, NULL, NULL, 6, NULL, 7, NULL, NULL), 9, true},
	{MakeTree(5, 3, 2, NULL, NULL, 4, NULL, NULL, 6, NULL, 7, NULL, NULL), 28, false},
}

func TestTwoSumIV(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findTarget(res.Input, res.Target)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
