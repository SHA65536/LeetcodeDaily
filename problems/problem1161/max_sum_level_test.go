package problem1161

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
	{MakeTree(1, 7, 7, NULL, NULL, -8, NULL, NULL, 0, NULL, NULL), 2},
	{MakeTree(989, NULL, 10250, 98693, NULL, NULL, -89388, NULL, -32127, NULL, NULL), 2},
}

func TestMaxLevelOfTree(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maxLevelSum(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
