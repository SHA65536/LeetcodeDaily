package problem2331

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
	{MakeTree(2, 1, NULL, NULL, 3, 0, NULL, NULL, 1, NULL, NULL), true},
	{MakeTree(1, NULL, NULL), true},
	{MakeTree(0, NULL, NULL), false},
}

func TestBooleanTreeEval(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := evaluateTree(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
