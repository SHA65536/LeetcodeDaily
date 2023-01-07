package problem0100

import (
	"fmt"
	"testing"

	. "leetcodedaily/helpers/binarytree"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Tree1, Tree2 *TreeNode
	Expected     bool
}

var Results = []Result{
	{MakeTree(1, 2, NULL, NULL, 3, NULL, NULL), MakeTree(1, 2, NULL, NULL, 3, NULL, NULL), true},
	{MakeTree(1, 2, NULL, NULL, NULL), MakeTree(1, NULL, 2, NULL, NULL), false},
	{MakeTree(1, 2, NULL, NULL, 1, NULL, NULL), MakeTree(1, 1, NULL, NULL, 2, NULL, NULL), false},
}

func TestIsSameTree(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := isSameTree(res.Tree1, res.Tree2)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
