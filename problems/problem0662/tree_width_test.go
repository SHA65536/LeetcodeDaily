package problem0662

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
	{MakeTree(1, 3, 5, NULL, NULL, 3, NULL, NULL, 2, NULL, 9, NULL, NULL), 4},
	{MakeTree(1, 3, 5, 6, NULL, NULL, NULL, NULL, 2, NULL, 9, 7, NULL, NULL, NULL), 7},
	{MakeTree(1, 3, 5, NULL, NULL, NULL, 2, NULL, NULL), 2},
}

func TestMaxWidthOfTree(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := widthOfBinaryTree(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
