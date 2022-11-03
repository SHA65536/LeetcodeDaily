package problem1038

import (
	"fmt"
	"testing"

	. "leetcodedaily/helpers/binarytree"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *TreeNode
	Expected *TreeNode
}

var Results = []Result{
	{MakeTree(4, 1, 0, NULL, NULL, 2, NULL, 3, NULL, NULL, 6, 5, NULL, NULL, 7, NULL, 8, NULL, NULL),
		MakeTree(30, 36, 36, NULL, NULL, 35, NULL, 33, NULL, NULL, 21, 26, NULL, NULL, 15, NULL, 8, NULL, NULL)},
	{MakeTree(0, NULL, 1, NULL, NULL),
		MakeTree(1, NULL, 1, NULL, NULL)},
}

func TestBstToGST(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := bstToGst(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
