package problem0234

import (
	"fmt"
	. "leetcodedaily/helpers/listnode"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    *ListNode
	Expected bool
}

var Results = []Result{
	{MakeListNode(1, 2, 2, 1), true},
	{MakeListNode(1, 2, 1), true},
	{MakeListNode(2, 2, 1), false},
	{MakeListNode(1, 1), true},
	{MakeListNode(1), true},
	{MakeListNode(1, 2), false},
}

func TestLinkedListPalindrome(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := isPalindrome(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
