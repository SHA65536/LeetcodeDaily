package problem0671

import (
	"fmt"
	"testing"

	. "leetcodedaily/helpers/binarytree"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *TreeNode
	Expected int
}

var Results = []Result{
	{MakeTree(2, 2, NULL, NULL, 5, 5, NULL, NULL, 7, NULL, NULL), 5},
	{MakeTree(2, 2, NULL, NULL, 2, NULL, NULL), -1},
}

func TestSecondMinimumNode(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findSecondMinimumValue(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
