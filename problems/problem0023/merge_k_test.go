package problem0023

import (
	"fmt"
	"testing"

	. "leetcodedaily/helpers/listnode"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []*ListNode
	Expected *ListNode
}

var Results = []Result{
	{
		[]*ListNode{
			MakeListNode(1, 4, 5),
			MakeListNode(1, 3, 4),
			MakeListNode(2, 6),
		},
		MakeListNode(1, 1, 2, 3, 4, 4, 5, 6),
	},
	{
		[]*ListNode{},
		MakeListNode(),
	},
	{
		[]*ListNode{
			MakeListNode(),
		},
		MakeListNode(),
	},
	{
		[]*ListNode{
			MakeListNode(1, 4, 5),
			MakeListNode(1, 3, 4),
			MakeListNode(),
			MakeListNode(2, 6),
		},
		MakeListNode(1, 1, 2, 3, 4, 4, 5, 6),
	},
}

func TestMergeKLists(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := mergeKLists(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
