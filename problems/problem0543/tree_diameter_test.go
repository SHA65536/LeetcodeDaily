package problem0543

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
	{MakeTree(1, 2, 4, NULL, NULL, 5, NULL, NULL, 3, NULL, NULL), 3},
	{MakeTree(1, 2, NULL, NULL, NULL), 1},
}

func TestTreeDiameter(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := diameterOfBinaryTree(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
