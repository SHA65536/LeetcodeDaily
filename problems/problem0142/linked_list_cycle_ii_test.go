package problem0142

import (
	. "leetcodedaily/helpers/listnode"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *ListNode
	Expected int
}

var Results = []Result{
	{MakeListNode(3, 2, 0, -4), 1},
	{MakeListNode(3, 2, 0, -4), 0},
	{MakeListNode(3, 2, 0, -4), 2},
	{MakeListNode(1, 2), 0},
	{MakeListNode(1), -1},
	{MakeListNode(), -1},
}

func TestLinkedListLoopII(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := makeLoop(res.Input, res.Expected)
		got := detectCycle(res.Input)
		assert.Equal(want, got)
	}
}

func makeLoop(l *ListNode, pos int) *ListNode {
	var cur, to *ListNode
	cur = l
	if pos == -1 {
		return nil
	}
	for i := 0; cur.Next != nil; i++ {
		if i == pos {
			to = cur
		}
		cur = cur.Next
	}
	cur.Next = to
	return to
}
