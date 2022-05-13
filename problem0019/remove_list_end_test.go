package problem0019

import (
	"fmt"
	. "leetcodedaily/helpers/listnode"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	List     *ListNode
	Num      int
	Expected *ListNode
}

var Results = []Result{
	{MakeListNode(1, 2, 3, 5, 6), 2, MakeListNode(1, 2, 3, 6)},
	{MakeListNode(1, 2, 3, 5, 6), 1, MakeListNode(1, 2, 3, 5)},
	{MakeListNode(1, 2, 3, 5, 6), 5, MakeListNode(2, 3, 5, 6)},
	{MakeListNode(1), 1, MakeListNode()},
	{MakeListNode(1, 2), 2, MakeListNode(2)},
}

func TestRemoveNthFromEnd(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := removeNthFromEnd(res.List, res.Num)
		assert.Equal(want, got, fmt.Sprintf("%d [%s]\nGot:  [%s]\nWant: [%s] ",
			res.Num, res.List.String(), got.String(), want.String()))
	}
}
