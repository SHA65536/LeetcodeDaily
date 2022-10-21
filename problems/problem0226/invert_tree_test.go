package problem0226

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
	{MakeTree(4, 2, 1, NULL, NULL, 3, NULL, NULL, 7, 6, NULL, NULL, 9, NULL, NULL),
		MakeTree(4, 7, 9, NULL, NULL, 6, NULL, NULL, 2, 3, NULL, NULL, 1, NULL, NULL)},
}

func TestInverTree(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := invertTree(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
