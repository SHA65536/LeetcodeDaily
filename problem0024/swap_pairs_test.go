package problem0024

import (
	"testing"

	. "leetcodedaily/helpers/listnode"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *ListNode
	Expected *ListNode
}

var Results = []Result{
	{MakeListNode(1, 2, 3, 4), MakeListNode(2, 1, 4, 3)},
	{MakeListNode(1, 2, 3), MakeListNode(2, 1, 3)},
	{MakeListNode(1), MakeListNode(1)},
	{MakeListNode(1, 2), MakeListNode(2, 1)},
	{MakeListNode(1, 2, 3, 4, 5, 6, 7, 8), MakeListNode(2, 1, 4, 3, 6, 5, 8, 7)},
	{MakeListNode(1, 2, 3, 4, 5, 6, 7), MakeListNode(2, 1, 4, 3, 6, 5, 7)},
}

func TestSwapPairs(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := swapPairs(res.Input)
		assert.Equal(want, got, got.String())
	}
}
