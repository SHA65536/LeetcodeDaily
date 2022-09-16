package problem0086

import (
	"fmt"
	. "leetcodedaily/helpers/listnode"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *ListNode
	X        int
	Expected *ListNode
}

var Results = []Result{
	{MakeListNode(1, 4, 3, 2, 5, 2), 3,
		MakeListNode(1, 2, 2, 4, 3, 5)},
	{MakeListNode(2, 1), 2,
		MakeListNode(1, 2)},
	{MakeListNode(1, 4, 3, 2, 5, 2), 1,
		MakeListNode(1, 4, 3, 2, 5, 2)},
	{MakeListNode(1, 4, 3, 2, 5, 2), 6,
		MakeListNode(1, 4, 3, 2, 5, 2)},
}

func TestPartition(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := partition(res.Input, res.X)
		assert.Equal(want, got, fmt.Sprintf("input:%s\nexpected:%s\nactual:%s", res.Input.String(), want.String(), got.String()))
	}
}
