package problem0106

import (
	"fmt"
	. "leetcodedaily/helpers/binarytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Inorder   []int
	Postorder []int
	Expected  *TreeNode
}

var Results = []Result{
	{
		[]int{1},
		[]int{1},
		MakeTree(1, NULL, NULL),
	},
	{
		[]int{9, 3, 15, 20, 7},
		[]int{9, 15, 7, 20, 3},
		MakeTree(3, 9, NULL, NULL, 20, 15, NULL, NULL, 7, NULL, NULL),
	},
}

func TestConstructTreeFromInPostOrder(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := buildTree(res.Inorder, res.Postorder)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
