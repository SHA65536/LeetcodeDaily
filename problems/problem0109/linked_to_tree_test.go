package problem0109

import (
	"fmt"
	. "leetcodedaily/helpers/binarytree"
	. "leetcodedaily/helpers/listnode"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *ListNode
	Expected *TreeNode
}

var Results = []Result{
	{
		MakeListNode(-10, -3, 0, 5, 9),
		MakeTree(0, -3, -10, NULL, NULL, NULL, 9, 5, NULL, NULL, NULL),
	},
}

func TestSortedListToBST(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := sortedListToBST(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
