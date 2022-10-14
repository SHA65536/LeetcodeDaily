package problem2095

import (
	"fmt"
	"testing"

	. "leetcodedaily/helpers/listnode"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *ListNode
	Expected *ListNode
}

var Results = []Result{
	{MakeListNode(1, 3, 4, 7, 1, 2, 6), MakeListNode(1, 3, 4, 1, 2, 6)},
	{MakeListNode(1, 2, 3, 4), MakeListNode(1, 2, 4)},
	{MakeListNode(2, 1), MakeListNode(2)},
	{MakeListNode(1), MakeListNode()},
	{MakeListNode(1, 2, 3), MakeListNode(1, 3)},
}

func TestDeleteMiddleNode(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := deleteMiddle(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
