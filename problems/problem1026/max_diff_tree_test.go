package problem1026

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
	{MakeTree(8, 3, 1, NULL, NULL, 6, 4, NULL, NULL, 7, NULL, NULL, 10, NULL, 14, 13, NULL, NULL, NULL), 7},
	{MakeTree(1, NULL, 2, NULL, 0, 3, NULL, NULL, NULL), 3},
}

func TestTreeMaxAncestorDiff(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maxAncestorDiff(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
