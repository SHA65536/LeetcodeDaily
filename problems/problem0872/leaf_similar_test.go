package problem0872

import (
	"fmt"
	"testing"

	. "leetcodedaily/helpers/binarytree"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	TreeA, TreeB *TreeNode
	Expected     bool
}

var Results = []Result{
	{
		MakeTree(3, 5, 6, NULL, NULL, 2, 7, NULL, NULL, 4, NULL, NULL, 1, 9, NULL, NULL, 8, NULL, NULL),
		MakeTree(3, 5, 6, NULL, NULL, 7, NULL, NULL, 1, 4, NULL, NULL, 2, 9, NULL, NULL, 8, NULL, NULL),
		true,
	},
	{
		MakeTree(1, 2, NULL, NULL, 3, NULL, NULL),
		MakeTree(1, 3, NULL, NULL, 2, NULL, NULL),
		false,
	},
}

func TestLeafSimilarTrees(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := leafSimilar(res.TreeA, res.TreeB)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
