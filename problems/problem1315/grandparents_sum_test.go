package problem1315

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
	{MakeTree(6, 7, 2, 9, NULL, NULL, NULL, 7, 1, NULL, NULL, 4, NULL, NULL, 8, 1, NULL, NULL, 3, NULL, 5, NULL, NULL), 18},
	{MakeTree(1), 0},
}

func TestEvenGrandparentSum(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := sumEvenGrandparent(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
