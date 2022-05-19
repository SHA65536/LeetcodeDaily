package problem0876

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
	{MakeListNode(1, 2, 3, 4, 5), MakeListNode(3, 4, 5)},
	{MakeListNode(1, 2, 3, 4, 5, 6), MakeListNode(4, 5, 6)},
	{MakeListNode(1), MakeListNode(1)},
	{MakeListNode(1, 2), MakeListNode(2)},
	{MakeListNode(1, 2, 3), MakeListNode(2, 3)},
	{MakeListNode(), MakeListNode()},
}

func TestMiddleNode(t *testing.T) {
	assert := assert.New(t)
	for _, res := range Results {
		want := res.Expected
		got := middleNode(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%s\n%s\n%s", res.Input, res.Expected, got))
	}
}
