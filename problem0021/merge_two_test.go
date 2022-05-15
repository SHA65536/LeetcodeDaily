package problem0021

import (
	"fmt"
	"testing"

	. "leetcodedaily/helpers/listnode"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	List1    *ListNode
	List2    *ListNode
	Expected *ListNode
}

var Results = []Result{
	{MakeListNode(1, 2, 4), MakeListNode(1, 2, 3), MakeListNode(1, 1, 2, 2, 3, 4)},
	/*{nil, nil, nil},
	{nil, MakeListNode(0), MakeListNode(0)},
	{MakeListNode(1, 2, 3), nil, MakeListNode(1, 2, 3)},*/
}

func TestMergeTwoLists(t *testing.T) {
	assert := assert.New(t)
	mergeTwoPrinter := func(r Result, g *ListNode) string {
		return fmt.Sprintf(
			"L1 - %s\nL2 - %s\nExp - %s\nGot - %s",
			r.List1.String(),
			r.List2.String(),
			r.Expected.String(),
			g.String())
	}

	for _, res := range Results {
		want := res.Expected
		got := mergeTwoLists(res.List1, res.List2)
		assert.Equal(want, got, mergeTwoPrinter(res, got))
	}
}
