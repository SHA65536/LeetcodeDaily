package problem0515

import (
	"fmt"
	. "leetcodedaily/helpers/binarytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *TreeNode
	Expected []int
}

var Results = []Result{
	{MakeTree(1, 3, 5, NULL, NULL, 3, NULL, NULL, 2, NULL, 9, NULL, NULL), []int{1, 3, 9}},
	{MakeTree(1, 2, NULL, NULL, 3, NULL, NULL), []int{1, 3}},
}

func TestLargestNodeLevel(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := largestValues(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
