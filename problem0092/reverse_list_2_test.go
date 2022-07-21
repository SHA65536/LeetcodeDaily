package problem0092

import (
	"fmt"
	. "leetcodedaily/helpers/listnode"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input       *ListNode
	Left, Right int
	Expected    *ListNode
}

var Results = []Result{
	{MakeListNode(1, 2, 3, 4, 5), 2, 4,
		MakeListNode(1, 4, 3, 2, 5)},
	{MakeListNode(5), 1, 1,
		MakeListNode(5)},
	{MakeListNode(1, 2, 3, 4, 5), 1, 5,
		MakeListNode(5, 4, 3, 2, 1)},
	{MakeListNode(1, 2, 3, 4, 5), 2, 5,
		MakeListNode(1, 5, 4, 3, 2)},
	{MakeListNode(1, 2, 3, 4, 5), 1, 4,
		MakeListNode(4, 3, 2, 1, 5)},
}

func TestNumMatchingSubseq(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := reverseBetween(res.Input, res.Left, res.Right)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
