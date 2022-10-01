package problem0572

import (
	"fmt"
	. "leetcodedaily/helpers/binarytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	T1, T2   *TreeNode
	Expected bool
}

var Results = []Result{
	{MakeTree(3, 4, 1, NULL, NULL, 2, NULL, NULL, 5, NULL, NULL), MakeTree(4, 1, NULL, NULL, 2, NULL, NULL), true},
	{MakeTree(3, 4, 1, NULL, NULL, 2, 0, NULL, NULL, NULL, 5, NULL, NULL), MakeTree(4, 1, NULL, NULL, 2, NULL, NULL), false},
}

func TestSkylineProblem(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := isSubtree(res.T1, res.T2)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
