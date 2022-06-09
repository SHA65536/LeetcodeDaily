package problem0206

import (
	. "leetcodedaily/helpers/listnode"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *ListNode
	Expected *ListNode
}

var Results = []Result{
	{
		MakeListNode(1, 2, 3, 4, 5),
		MakeListNode(5, 4, 3, 2, 1),
	},
	{
		MakeListNode(),
		MakeListNode(),
	},
	{
		MakeListNode(1),
		MakeListNode(1),
	},
	{
		MakeListNode(1, 2, 3, 4, 5, 6),
		MakeListNode(6, 5, 4, 3, 2, 1),
	},
}

func TestReverseList(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := reverseList(res.Input)
		assert.Equal(want, got, got.String())
	}
}
