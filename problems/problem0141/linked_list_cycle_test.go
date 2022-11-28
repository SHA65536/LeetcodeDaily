package problem0141

import (
	. "leetcodedaily/helpers/listnode"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *ListNode
	Pos      int
	Expected bool
}

var Results = []Result{
	{MakeListNode(3, 2, 0, -4), 1, true},
	{MakeListNode(1, 2), 0, true},
	{MakeListNode(1), -1, false},
	{MakeListNode(3, 2, 0, -4), -1, false},
}

func TestLinkedListLoopI(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		makeLoop(res.Input, res.Pos)
		got := hasCycle(res.Input)
		assert.Equal(want, got)
	}
}

func makeLoop(l *ListNode, pos int) {
	var cur, to *ListNode
	cur = l
	if pos == -1 {
		return
	}
	for i := 0; cur.Next != nil; i++ {
		if i == pos {
			to = cur
		}
		cur = cur.Next
	}
	cur.Next = to
}
