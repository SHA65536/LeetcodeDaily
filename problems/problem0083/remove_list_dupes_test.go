package problem0083

import (
	"fmt"
	. "leetcodedaily/helpers/listnode"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *ListNode
	Expected *ListNode
}

var Results = []Result{
	{MakeListNode(1, 1, 2), MakeListNode(1, 2)},
	{MakeListNode(1, 1, 2, 3, 3), MakeListNode(1, 2, 3)},
	{MakeListNode(1), MakeListNode(1)},
	{MakeListNode(), MakeListNode()},
	{MakeListNode(1, 2), MakeListNode(1, 2)},
}

func TestRemoveListDuplicates(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := deleteDuplicates(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
