package problem0501

import (
	"fmt"
	. "leetcodedaily/helpers/binarytree"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *TreeNode
	Expected []int
}

var Results = []Result{
	{MakeTree(1, NULL, 2, 2, NULL, NULL, NULL), []int{2}},
	{MakeTree(0), []int{0}},
	{MakeTree(2, 1, 1, NULL, NULL, 1, NULL, NULL, 3, 3, NULL, NULL, 3, NULL, NULL), []int{1, 3}},
}

func TestFindModesInTree(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findMode(res.Input)
		sort.Ints(got)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
