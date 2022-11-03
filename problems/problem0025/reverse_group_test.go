package problem0025

import (
	"fmt"
	"testing"

	. "leetcodedaily/helpers/listnode"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *ListNode
	K        int
	Expected *ListNode
}

var Results = []Result{
	{MakeListNode(1, 2, 3, 4, 5), 2, MakeListNode(2, 1, 4, 3, 5)},
	{MakeListNode(1, 2, 3, 4, 5), 3, MakeListNode(3, 2, 1, 4, 5)},
}

func TestReverseListInKGroups(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := reverseKGroup(res.Input, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
