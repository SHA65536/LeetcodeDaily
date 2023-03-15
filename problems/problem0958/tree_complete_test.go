package problem0958

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
	{MakeTree(1, 2, 4, NULL, NULL, 5, NULL, NULL, 3, 6, NULL, NULL, NULL), true},
	{MakeTree(1, 2, 4, 7, NULL, NULL, NULL, 5, NULL, NULL, 3, 6, NULL, NULL, 8, NULL, NULL), true},
	{MakeTree(1, 2, 4, NULL, NULL, 5, NULL, NULL, 3, NULL, 7, NULL, NULL), false},
}

func TestCheckTreeCompleteness(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := isCompleteTree(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
