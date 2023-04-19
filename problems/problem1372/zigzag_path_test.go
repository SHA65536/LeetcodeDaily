package problem1372

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
	{MakeTree(1, NULL, 1, 1, NULL, NULL, 1, 1, NULL, 1, NULL, 1, NULL, NULL, 1, NULL, NULL), 3},
	{MakeTree(1, 1, NULL, 1, 1, NULL, 1, NULL, NULL, 1, NULL, NULL, 1, NULL, NULL), 4},
	{MakeTree(1, NULL, NULL), 0},
}

func TestLongestZigZagPath(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := longestZigZag(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
