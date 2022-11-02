package problem0951

import (
	"fmt"
	"testing"

	. "leetcodedaily/helpers/binarytree"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *TreeNode
	Target   *TreeNode
	Expected bool
}

var Results = []Result{
	{
		MakeTree(1, 2, 4, NULL, NULL, 5, 7, NULL, NULL, 8, NULL, NULL, 3, 6, NULL, NULL, NULL),
		MakeTree(1, 3, NULL, 6, NULL, NULL, 2, 4, NULL, NULL, 5, 8, NULL, NULL, 7, NULL, NULL),
		true,
	},
	{
		MakeTree(),
		MakeTree(),
		true,
	},
	{
		MakeTree(),
		MakeTree(1),
		false,
	},
	{
		MakeTree(1, 2, 4, NULL, NULL, 8, 7, NULL, NULL, 5, NULL, NULL, 3, 6, NULL, NULL, NULL),
		MakeTree(1, 3, NULL, 6, NULL, NULL, 2, 4, NULL, NULL, 5, 8, NULL, NULL, 7, NULL, NULL),
		false,
	},
}

func TestCanFlipEqualTree(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := flipEquiv(res.Input, res.Target)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
