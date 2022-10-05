package problem1609

import (
	"fmt"
	. "leetcodedaily/helpers/binarytree"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *TreeNode
	Expected bool
}

var Results = []Result{
	{MakeTree(1, 10, 3, 12, NULL, NULL, 8, NULL, NULL, NULL, 4, 7, 6, NULL, NULL, NULL, 9, NULL, 2, NULL, NULL), true},
	{MakeTree(5, 4, 3, NULL, NULL, 3, NULL, NULL, 2, 7, NULL, NULL, NULL), false},
	{MakeTree(5, 9, 3, NULL, NULL, 5, NULL, NULL, 1, 7, NULL, NULL, NULL), false},
}

func TestEvenOddTree(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := isEvenOddTree(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
