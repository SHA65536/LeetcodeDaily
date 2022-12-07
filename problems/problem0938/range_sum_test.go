package problem0938

import (
	"fmt"
	. "leetcodedaily/helpers/binarytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input     *TreeNode
	Low, High int
	Expected  int
}

var Results = []Result{
	{
		MakeTree(10, 5, 3, NULL, NULL, 7, NULL, NULL, 15, NULL, 18, NULL, NULL),
		7, 15, 32,
	},
	{
		MakeTree(10, 5, 3, 1, NULL, NULL, NULL, 7, 6, NULL, NULL, NULL, 15, 13, NULL, NULL, 18, NULL, NULL),
		6, 10, 23,
	},
}

func TestRangeSumBST(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := rangeSumBST(res.Input, res.Low, res.High)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
