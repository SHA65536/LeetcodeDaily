package problem0652

import (
	"fmt"
	. "leetcodedaily/helpers/binarytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *TreeNode
	Expected []*TreeNode
}

var Results = []Result{
	{
		MakeTree(1, 2, 4, NULL, NULL, NULL, 3, 2, 4, NULL, NULL, NULL, 4, NULL, NULL),
		[]*TreeNode{
			MakeTree(4, NULL, NULL),
			MakeTree(2, 4, NULL, NULL, NULL),
		},
	},
	{
		MakeTree(2, 1, NULL, NULL, 1, NULL, NULL),
		[]*TreeNode{
			MakeTree(1, NULL, NULL),
		},
	},
	{
		MakeTree(2, 2, 3, NULL, NULL, NULL, 2, 3, NULL, NULL, NULL),
		[]*TreeNode{
			MakeTree(3, NULL, NULL),
			MakeTree(2, 3, NULL, NULL, NULL),
		},
	},
	{
		MakeTree(2, 1, 11, NULL, NULL, NULL, 11, 1, NULL, NULL, NULL),
		nil,
	},
}

func TestFindDuplicateTrees(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findDuplicateSubtrees(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
