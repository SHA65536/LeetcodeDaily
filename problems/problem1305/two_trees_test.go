package problem1305

import (
	"fmt"
	. "leetcodedaily/helpers/binarytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Root1, Root2 *TreeNode
	Expected     []int
}

var Results = []Result{
	{
		MakeTree(2, 1, NULL, NULL, 4, NULL, NULL),
		MakeTree(1, 0, NULL, NULL, 3, NULL, NULL),
		[]int{0, 1, 1, 2, 3, 4},
	},
	{
		MakeTree(1, NULL, 8, NULL, NULL),
		MakeTree(8, 1, NULL, NULL, NULL),
		[]int{1, 1, 8, 8},
	},
	{
		MakeTree(1, NULL, 8, NULL, NULL),
		MakeTree(2, 1, NULL, NULL, 4, NULL, NULL),
		[]int{1, 1, 2, 4, 8},
	},
}

func TestTwoSearchTrees(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := getAllElements(res.Root1, res.Root2)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
