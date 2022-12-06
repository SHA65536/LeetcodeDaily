package problem0328

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
	{MakeListNode(1, 2, 3, 4, 5), MakeListNode(1, 3, 5, 2, 4)},
	{MakeListNode(2, 1, 3, 5, 6, 4, 7), MakeListNode(2, 3, 6, 7, 1, 5, 4)},
	{MakeListNode(), MakeListNode()},
}

func TestEvenOddList(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := oddEvenList(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
