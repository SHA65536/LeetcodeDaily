package problem0101

import (
	"fmt"
	"testing"

	. "leetcodedaily/helpers/binarytree"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *TreeNode
	Expected bool
}

var Results = []Result{
	{
		MakeTree(1, 2, 3, NULL, NULL, 4, NULL, NULL, 2, 4, NULL, NULL, 3, NULL, NULL),
		true,
	},
	{
		MakeTree(1, 2, NULL, 3, NULL, NULL, 2, NULL, 3, NULL, NULL),
		false,
	},
}

func TestIsSymetricTree(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := isSymmetric(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
