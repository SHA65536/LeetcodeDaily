package problem1339

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
	{MakeTree(1, 2, 4, NULL, NULL, 5, NULL, NULL, 3, 6, NULL, NULL, NULL), 110},
	{MakeTree(1, NULL, 2, 3, NULL, NULL, 4, 5, NULL, NULL, 6, NULL, NULL), 90},
}

func TestMaxSplitTree(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maxProduct(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
